package main

import (
	"douyin-template/services/service_user/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.Static("/static", "./public")
	preGroup := engine.Group("/douyin")

	// basic apis
	preGroup.GET("/user/", controller.UserInfo)
	preGroup.POST("/user/register/", controller.Register)
	preGroup.POST("/user/login/", controller.Login)

	preGroup.GET("/feed/", controller.Feed)
	preGroup.POST("/publish/action/", controller.Publish)
	preGroup.GET("/publish/list/", controller.PublishList)
}
