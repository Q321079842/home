package router

import (
	"github.com/gin-gonic/gin"
	. "user/controller"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.NoRoute(NotFound)
	user := engine.Group("user")
	{
		user.POST("/add", Add)
		user.GET("/delete", Delete)
		user.PUT("/edit", Edit)
		user.GET("/find", Find)
	}
	engine.Run(":8080")

}
