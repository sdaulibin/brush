package model

type Base struct {
	Id int64 `json:"id" gorm:"primary_key"`
	//CreateAt *time.Time `json:"-"`
	//UpdateAt *time.Time `json:"-"`
}
