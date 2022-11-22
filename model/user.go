package model

func (user User) TableName() string {
	return "user"
}

type UserInfo struct {
	Token    string `json:"token"`
	UserName string `json:"userName"`
}

type User struct {
	Base
	Name  string
	Score int
	Token string
}
