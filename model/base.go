package model

import "binginx.com/brush/internal/logs"

type Base struct {
	Id int64 `json:"id" gorm:"primary_key"`
	//CreateAt *time.Time `json:"-"`
	//UpdateAt *time.Time `json:"-"`
}

func LogErr(errType string, err error) {
	logs.Logger.Error(errType, err)
}
