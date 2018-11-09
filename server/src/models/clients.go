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
	ID          uint      `json:"id"`
	Pseudo      string    `json:"pseudo"`
	Password    string    `json:"password"`
	Ips         []*Ip     `gorm:"many2many:client_address";json:"ips"`
	Friends     []*Client `sql:"-"`
	Friendships []*Friendship
	Token       string `json:"token";sql:"-"`
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
	err = client.Preload(true)
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

func (client *Client) getFriendship() ([]*Friendship, error) {
	var friendships []*Friendship
	err := GetDB().Table("friendships").Where("client_id = ?", client.ID).Find(&friendships).Error
	return friendships, err
}

func (client *Client) Preload(friend bool) error {
	err := GetDB().Preload("Ips").First(&client).Error
	if err != nil {
		return err
	}
	client.Friendships, err = client.getFriendship()
	if err != nil {
		return err
	}
	if friend {
		client.getFriends(client.Friendships)
	}
	return err
}

func (client *Client) getFriends(friendships []*Friendship) {
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
	client.Preload(true)
	if err != nil {
		return nil, err
	}
	if client.Pseudo == "" { //User not found!
		return nil, errors.New("Pseudo is empty")
	}

	client.Password = ""
	return client, nil
}

func (client *Client) Update() (*Client, error) {
	//check if ip has changed
	if len(client.Ips) > 0 {
		err := GetDB().Model(&client).Association("Ips").Replace(client.Ips).Error
		if err != nil {
			return nil, err
		}
	}
	//if len(client.Friends) != len(client.Friendships) {	//@TODO find a way to know which whether a friend was add or already there
	//	var friendships []*Friendship
	//	for i:=0; i < len(client.Friends); i ++ {
	//		friendships = append(friendships, &Friendship{
	//			FriendID: client.Friends[i].ID,
	//			ClientID:client.ID,
	//			Accepted:false,
	//		})
	//	}
	//	client.Friendships = friendships
	//}
	error := GetDB().Save(&client).Error
	return client, error
}

func (client *Client) Delete() (map[string]interface{}) {
	response := u.Message(true, "Client has been deleted", http.StatusOK)
	response["client"] = client
	return response
}

func (client *Client) AddFriend(friend *Client) (*Client, error) {
	if friend.ID == 0 {
		_friend, err := GetClientFromPseudo(friend)
		friend = _friend
		if err != nil {
			return nil, err
		}
	}
	client.Friends = append(client.Friends, friend)
	client.addFriendShip(friend)
	client, err := client.Update()
	return client, err
}

func (client *Client) RemoveFriend(friend *Client) {
	friends := client.Friends
	indice := -1
	for i := 0; i < len(friends); i ++ {
		if friend.ID == friends[i].ID {
			indice = i
		}
	}
	if indice == -1 {
		return
	}
	friends[indice] = friends[len(friends)-1]
	friends = friends[:len(friends)-1]
	client.Friends = friends
	client.Update()
}
func (client *Client) RegisterFriends() {
	client.Friendships = []*Friendship{}
	for i := 0; i < len(client.Friends); i++ {
		client.addFriendShip(client.Friends[i])
	}
}

func (client *Client) addFriendShip(friend *Client) {
	accepted := false
	log.Print("addFriendship call getClient")
	friendships, err := friend.getFriendship()
	if err != nil {
		log.Printf("Trouble in client::addFriendship: %v", err)
	}
	for i := 0; i < len(friendships); i ++ {
		log.Printf("On est dans la boucle for: %d", i)
		if friendships[i].FriendID == client.ID {
			accepted = true
		}
		friendshipUpdate := friendships[i]
		friendshipUpdate.Accepted = accepted
		GetDB().Save(&friendshipUpdate)
	}
	friendship := &Friendship{
		FriendID: friend.ID,
		ClientID: client.ID,
		Accepted: accepted,
	}
	client.Friendships = append(client.Friendships, friendship)
}
