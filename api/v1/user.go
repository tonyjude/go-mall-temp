package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall-temp/service"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// UserLoginHandler 用户登陆接口
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
