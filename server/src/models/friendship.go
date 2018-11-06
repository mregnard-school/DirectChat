package models

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
	return GetClient(friendship.FriendID)
}