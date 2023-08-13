package service

import (
	"context"
	"go-mall-temp/dao"
	"go-mall-temp/model"
	"go-mall-temp/pkg/e"
	"go-mall-temp/pkg/util"
	"go-mall-temp/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name" `
	UserName string `json:"user_name" form:"user_name" `
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.SUCCESS

	if service.Key == "" || len(service.Key) != 16 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "秘钥长度不足",
		}
	}

	if len(service.UserName) < 5 || len(service.UserName) > 20 {
		code = e.ErrorUserNameLENGTHERR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "用户名长度5-20之间",
		}
	}

	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.jpg",
		Money:    util.Encrypt.AesEncoding("10000"),
	}

	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.ERROR
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// UserLogin 用户登陆函数
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist { // 如果查询不到，返回相应的错误
		//log.LogrusObj.Error(err)
		code := e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册！",
		}
	}

	if !user.CheckPassword(service.Password) {
		code := e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}

	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		//log.LogrusObj.Error(err)
		code := e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录！",
		}
	}

	return serializer.Response{
		Status: e.SUCCESS,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(e.SUCCESS),
	}
}
