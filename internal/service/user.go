package service

import (
	"binginx.com/brush/internal/clients"
	"binginx.com/brush/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) error {
	return clients.WriteDBCli.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			model.LogErr("CreateUser from WriteDBCli db", err)
			return err
		}
		return nil
	})
}

func GetUserByPhone(phone string) (user *model.User, err error) {
	err = clients.ReadDBCli.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		model.LogErr("GetUser from ReadDBCli db", err)
		return nil, err
	}
	return user, nil
}

func GetUserByToken(token string) (user *model.User, err error) {
	err = clients.ReadDBCli.Where("token = ?", token).First(&user).Error
	if err != nil {
		model.LogErr("GetUser from ReadDBCli db", err)
		return nil, err
	}
	return user, nil
}

func UpdateUserScore(phone, score string) error {
	if err := clients.WriteDBCli.Table(model.User{}.TableName()).
		Where("phone = ?", phone).
		Update("score", score).Error; err != nil {
		model.LogErr("UpdateUser from WriteDBCli db", err)
		return err
	}
	return nil
}

func UpdateUser(user *model.User) error {
	if err := clients.WriteDBCli.Save(user).Error; err != nil {
		model.LogErr("UpdateUser from WriteDBCli db", err)
		return err
	}
	return nil
}
