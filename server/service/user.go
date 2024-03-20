package service

import (
	"snake/global"
	"snake/model"
)

type UserService struct {
}

func (u *UserService) Login(userName string, password string) bool {
	var user model.SysUser
	global.DB.Where("user_name = ?", userName).First(&user)
	if user.Password == password {
		return true
	}
	return false
}
