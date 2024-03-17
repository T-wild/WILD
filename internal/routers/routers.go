package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"wild/internal/controllers"
	"wild/internal/middlewares"
	"wild/internal/pkg/Logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 连通测试
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("server.version"))
	})

	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登陆业务路由
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		v1.GET("/user", controllers.GetUsersHandler)
	}

	return r
}
