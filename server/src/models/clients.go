package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	u "server/utils"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}
//  Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;foreignkey:idPerson;association_foreignkey:idAccount;association_jointable_foreignkey:account_id;jointable_foreignkey:person_id;"`
//a struct to rep user client
type Client struct {
	ID 			uint	 		`json:"id"`
	Pseudo   	string   		`json:"pseudo"`
	Password 	string   		`json:"password"`
	Ips      	[]*Ip     		`gorm:"many2many:client_address";json:"ips"`
	Friends  	[]*Client  		`sql:"-"`
	Friendships	[]*Friendship
	Token    	string   		`json:"token";sql:"-"`
}

func (*Client) TableName() string {
	return "clients"
}

func (client *Client) Create() (*Client, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
	client.Password = string(hashedPassword)
	client.RegisterFriends()
	GetDB().Create(client)

	if client.ID <= 0 {

		return nil, errors.New("Failed to create client, connection error.")
	}

	//Create new JWT token for the newly registered client
	tk := &Token{UserId: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(hashedPassword))
	client.Token = tokenString
	//for test it sucked
	client.Password = "" //delete password
	client.setEmptyValues()
	return client, nil
}

func (client *Client) setEmptyValues() {
	if client.Friends == nil {
		client.Friends = []*Client{}
	}
	if client.Friendships == nil {
		client.Friendships = []*Friendship{}
	}
}
func GetClient(u uint) (*Client, error) {

	client := &Client{}
	err := GetDB().Table("clients").Where("id = ?", u).First(client).Error
	if err != nil {
		return nil, err
	}
	err = client.Preload()
	if err != nil {
		return nil, err
	}
	if client.Pseudo == "" { //User not found!
		return nil, errors.New("Pseudo is empty")
	}
	client.Password = ""
	client.setEmptyValues()
	return client, nil
}

func (client *Client) Preload() error {
	err := GetDB().Preload("Ips").First(&client).Error
	if err != nil {
		return err
	}
	var friendships []*Friendship
	client.getFriends(friendships)
	return err
}

func (client *Client) getFriends(friendships []*Friendship) {
	GetDB().Table("friendships").Where("client_id = ?", client.ID).Find(&friendships)
	var friends []*Client
	for i := 0; i < len(friendships); i++ {
		friend, err := friendships[i].getFriend()
		if err != nil {
			log.Printf("Error loading friends: %s", err.Error())
		}
		friends = append(friends, friend)
	}
	client.Friends = friends
}

func GetClientFromPseudo(friend *Client) (*Client, error) {
	pseudo := friend.Pseudo
	client := &Client{}
	err := GetDB().Table("clients").Where("pseudo = ?", pseudo).First(client).Error
	GetDB().Preload("Ips").First(&client)
	GetDB().Preload("Friends").First(&client)
	if err != nil {
		return nil, err
	}
	if client.Pseudo == "" { //User not found!
		return nil, errors.New("Pseudo is empty")
	}

	client.Password = ""
	return client, nil
}

func (client *Client) Update() (map[string] interface{})  {
	//check if ip has changed
	if len(client.Ips) > 0 {
		GetDB().Model(&client).Association("Ips").Replace(client.Ips)
	}
	GetDB().Save(&client)

	resp := u.Message(true, "Client updated", http.StatusOK)
	resp["client"] = client
	return resp
}

func (client *Client) Delete() (map[string] interface{}) {
	response := u.Message(true, "Client has been deleted", http.StatusOK)
	response["client"] = client
	return response
}

func (client *Client) AddFriend(friend *Client) (map[string] interface{}){
	response := u.Message(true, "Client has been created", http.StatusOK)
	client.Friends = append(client.Friends, friend)
	client.addFriendShip(friend)
	client.Update()
	response["client"] = client
	return response
}

func (client *Client) RemoveFriend(friend *Client) {
	friends := client.Friends
	indice := -1
	for i:=0; i < len(friends); i ++ {
		if friend.ID == friends[i].ID {
			indice = i
		}
	}
	if indice == -1 {
		return
	}
	friends[indice] = friends[len(friends) - 1]
	friends = friends[:len(friends) - 1]
	client.Friends = friends
	client.Update()
}
func (client *Client) RegisterFriends() {
	var friendships []*Friendship
	for i:=0; i< len(client.Friends); i++ {
		friendship := &Friendship{
			FriendID:client.Friends[i].ID,
			ClientID:client.ID,
			Accepted:false,
		}
		friendships = append(friendships, friendship)
	}
	client.Friendships = friendships
	//GetDB().Save(friendships)
}

func (client *Client) addFriendShip(friend *Client)  {
	friendship := &Friendship{
		FriendID:friend.ID,
		ClientID:client.ID,
		Accepted:false,
	}
	client.Friendships = append(client.Friendships, friendship)
}
