package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"server/models"
	u "server/utils"
)

func Login(pseudo string, password string, _ip string) (*models.Client, int, string) {

	client := &models.Client{}
	err := models.GetDB().Table("clients").Where("pseudo = ?", pseudo).First(client).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusUnauthorized, "Client doesn't exists"
		}
		return nil, http.StatusInternalServerError, "Server Error"
	}

	client.Preload()
	if _ip != "" {
		ip := &models.Ip{
			Address: _ip,
		}
		client.Ips = []*models.Ip{ip}
	}
	client.Update()
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

//Validate incoming user details...
func Validate(client *models.Client) (map[string] interface{}, bool) {

	if len(client.Pseudo) < 1 {
		return u.Message(false, "Pseudo is required", http.StatusUnprocessableEntity), false
	}

	if len(client.Password) < 6 {
		return u.Message(false, "Password is required and need at least 6 characters", http.StatusUnprocessableEntity), false
	}

	//Pseudo must be unique
	temp := &models.Client{}

	//check for errors and duplicate pseudos
	err := models.GetDB().Table("clients").Where("pseudo = ?", client.Pseudo).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry", http.StatusInternalServerError), false
	}
	if temp.Pseudo != "" {
		return u.Message(false, "Pseudo address already in use by another user.", http.StatusUnprocessableEntity), false
	}

	return u.Message(false, "Requirement passed", http.StatusOK), true
}