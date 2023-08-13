package routes

import (
	"github.com/gin-gonic/gin"
	api "go-mall-temp/api/v1"
	"go-mall-temp/middleware"
	"net/http"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
		}
	}

	return r
}
