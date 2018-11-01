package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
func Open(username string, password string, dbName string, dbHost string) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, dbHost,  dbName)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
		log.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Client{}, &Ip{})
}

/**
Jointable ForeignKey
If you want to change join table’s foreign keys, you could use tag association_jointable_foreignkey, jointable_foreignkey

type CustomizePerson struct {
  IdPerson string             `gorm:"primary_key:true"`
}

type CustomizeAccount struct {
  IdAccount string `gorm:"primary_key:true"`
  Name      string
}
 */

func GetDB() *gorm.DB {
	return db
}
