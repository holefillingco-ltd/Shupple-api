package server

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	uGroup := r.Group("/users")
	{
		ctrl := controller.UserController{}
		uGroup.GET("", ctrl.Shupple)
		uGroup.POST("", ctrl.Create)
		uGroup.PUT("", ctrl.Update)
		uGroup.GET("/select", ctrl.GetUser)
		uGroup.POST("/compatible", ctrl.CreateCompatible)
	}

	return r
}
