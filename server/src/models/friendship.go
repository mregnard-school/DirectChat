package models

import "errors"

type Friendship struct {
	ID 			uint
	ClientID	uint
	FriendID	uint
	Accepted 	bool
}

func (*Friendship) TableName() string {
	return "friendships"
}

func (friendship *Friendship) getFriend() (*Client, error) {

	return GetClientWithoutFriend(friendship.FriendID)
}
func GetClientWithoutFriend(u uint) (*Client, error) {
	client := &Client{}
	err := GetDB().Table("clients").Where("id = ?", u).First(client).Error
	if err != nil {
		return nil, err
	}
	err = client.Preload(false)
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
