package logic

import (
	"chat/helper"
	"chat/internal/types"
	"chat/models"
	"errors"
	"log"
)

func UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {

	//判断用户名是否一致
	count, err := models.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("用户名已存在")
		return
	}
	//数据入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	n, err := models.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user rouw:", n)
	return
}
