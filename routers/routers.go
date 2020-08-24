package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zhouwy1994/gin-example/api/v1"
	"github.com/zhouwy1994/gin-example/middleware/jwt"
	"github.com/zhouwy1994/gin-example/pkg/setting"
)

func registerRouters(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", v1.UserRegister)
		apiv1.POST("/login", v1.UserLogin)
		apiv1.Use(jwt.JWT())
	}
}

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()
	// 注册日志gin日志中间件
	r.Use(gin.Logger())
	// 注册gin错误恢复器中间件
	r.Use(gin.Recovery())

	// 注册路由信息
	registerRouters(r)

	return r
}
