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

type JicUser struct {
	Code        int         `json:"code"`
	Msg         string      `json:"msg"`
	JicUserInfo JicUserInfo `json:"data"`
}

type JicUserInfo struct {
	EmployeeID int    `json:"employeeID,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
}
