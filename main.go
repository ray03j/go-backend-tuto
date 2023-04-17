// 1→
package main

import (
	// "github.com/gin-gonic/gin"
	
	"go-backend-tuto/db"
	// pathが通ってる
	"go-backend-tuto/server"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Hello world")
	// })

	db.Init()
	// router.Run()
  
  server.Init()
	db.Close()
}

/* 1→docker-compose.yml

package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World")
    })
        r.Run()
}

サーバを起動して、エンドポイントを叩く
$ go run main.go

$ curl localhost:8080
*/