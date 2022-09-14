package model

type Contents struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []Content `json:"data"`
}

type Content struct {
	ContentId int `json:"contentId"`
}
