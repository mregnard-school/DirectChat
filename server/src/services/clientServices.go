package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"server/models"
)

func Login(pseudo string, password string) (*models.Client, int, string) {

	client := &models.Client{}
	err := models.GetDB().Table("clients").Where("pseudo = ?", pseudo).First(client).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusUnauthorized, "Client doesn't exists"
		}
		return nil, http.StatusInternalServerError, "Server Error"
	}

	client.Preload()

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, http.StatusUnauthorized, "Wrong password"
	}
	//Worked! Logged In
	client.Password = ""

	//Create JWT token
	tk := &models.Token{UserId: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	client.Token = tokenString //Store the token in the response

	return client, http.StatusOK, "Logged In"
}