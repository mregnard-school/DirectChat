package models

type Friendship struct {
	ID			uint
	Client 		Client	`gorm:"foreignkey:ID"`
	Friend 		Client	`gorm:"foreignkey:ID"`
	Accepted 	bool
}
