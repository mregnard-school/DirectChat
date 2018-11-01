package models

import (
	"github.com/jinzhu/gorm"
	u "server/utils"
)

type Ip struct {
	gorm.Model
	Address string `json:"address"`
}

func (address *Ip) Validate() (map[string] interface{}, bool){

	//temp := &Address{}
	//err := GetDB().Table("clients").Where("id = ?", address.Address).First(temp).Error
	//
	//if err != nil {
	//	return u.Message(false, "The client doesn't exist"), false
	//}
	return u.Message(false, "Requirement passed"), true

}
func (address *Ip) Create() (map[string] interface{}) {
	if resp, ok := address.Validate(); !ok {
		return resp
	}

	GetDB().Create(address)

	if address.ID <= 0 {
		return u.Message(false, "failed to create address, connection error")
	}
	response := u.Message(true, "Address has been created")
	response["address"] = address
	return response
}

func GetIp(u uint) *Ip {
	address := &Ip{}
	GetDB().Table("addresss").Where("id = ?", u).First(address)
	if address.Address == "" {
		return nil
	}
	return address
}

