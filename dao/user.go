package dao

import (
	"lss/config"
	"lss/model"
)

var User = &user{}

type user struct{}

func (user) Info(id int) (*model.User, error) {
	var info *model.User
	conn := config.DBObj.GetMysqlConn()
	if conn == nil {
		return nil, conn.Error
	}
	if err := conn.Where("id = ?", id).Find(&info).Error; err != nil {
		return nil, err
	}

	return info, nil
}
