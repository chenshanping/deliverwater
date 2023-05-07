package service

import (
	"deliverwater/global"
	"deliverwater/model"
)

var UserService = userService{}

type userService struct {
}

func (usrv userService) UserList(pageSize int, offsetval int) (userList []*model.UserInfo, total int64) {
	global.DB.Model(userList).Count(&total).Limit(pageSize).Offset(offsetval).Find(&userList)
	return
}
