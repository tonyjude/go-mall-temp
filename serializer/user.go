package serializer

import "go-mall-temp/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NikeName string `json:"nike_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NikeName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
