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

	u := r.Group("/users")
	{
		ctrl := controller.UserController{}
		// テスト用 確認取れれば消して良い
		u.GET("", ctrl.Index)
		u.POST("", ctrl.Create)
	}

	return r
}
