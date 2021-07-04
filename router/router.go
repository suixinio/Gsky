package router

import (
	v1 "Gsky/api/v1"
	"Gsky/middleware"
	"Gsky/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("upload/", v1.Upload)
		// 返回index.html
		routerV1.GET("/", func(c *gin.Context) {

		})
		// 登陆返回带有token的页面
		routerV1.POST("/", func(c *gin.Context) {

		})
	}

	r.Run(utils.HttpPort)
}
