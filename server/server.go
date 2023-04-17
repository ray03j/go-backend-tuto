// routingの実装
package server

import (
	"github.com/gin-gonic/gin"

	"go-backend-tuto/controller"
)

// Init is initialize server
func Init() {
	r := router() //router関数は後で定義
	r.Run()
}

func router() *gin.Engine {
	router := gin.Default()

	usr := router.Group("/users")
	{
		// どのpathのとき、どの処理をするか
		ctrl := user.Controller{}
		usr.GET("",ctrl.Index)
		usr.GET("/:id", ctrl.Show)
		usr.POST("", ctrl.Create)
		usr.PUT("/:id", ctrl.Update)
		usr.DELETE("/:id", ctrl.Delete)
	}

	return router
}


