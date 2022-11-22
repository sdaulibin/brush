package service

import (
	"binginx.com/brush/internal/clients"
	"binginx.com/brush/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) error {
	return clients.WriteDBCli.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetUser(token string) (user model.User, err error) {
	err = clients.ReadDBCli.Where("token = ?", token).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
