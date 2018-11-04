package models

type Friendship struct {
	ID			uint
	Client 		*Client
	Friend 		*Client		`gorm:"association_foreignkey:ID"`
	Accepted 	bool
}

func (*Friendship) TableName() string {
	return "friendships"
}