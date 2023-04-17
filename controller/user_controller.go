// アクションの追加

// AbortWithStatusとは
// gin.Hはmap[string]interface{}と同義。
// map[string]interface{}とは
// どこまでの処理をしてるか
package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-backend-tuto/service"
)

// Controller is user controlller
type Controller struct{}



// Index action: GET /users
func (pc Controller) Index(ctx *gin.Context) {
	var s user.Service
	p, err := s.GetAll()


	if err != nil {
		ctx.AbortWithStatus(404) // Abort with :中止する
		fmt.Println(err)
	} else {
		ctx.JSON(200,p)
	}
}



// Create action: POST /users
func (pc Controller) Create(ctx *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(ctx)

	if err != nil {
		ctx.AbortWithStatus(400) // Abort with :中止する
		fmt.Println(err)
	} else {
		ctx.JSON(201,p)
	}
}



// Show action: GET /users/:id
func (pc Controller) Show(ctx *gin.Context) {
	id := ctx.Params.ByName("id") // ByNameメソッドは、与えられたKey(ここではid)からデータを取得する便利メソッド
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		ctx.AbortWithStatus(404) // Abort with :中止する
		fmt.Println(err)
	} else {
		ctx.JSON(200,p)
	}
}



// Update action: PUT /users/:id
func (pc Controller) Update(ctx *gin.Context) {
	id := ctx.Params.ByName("id") // ByNameメソッドは、与えられたKey(ここではid)からデータを取得する便利メソッド
	var s user.Service
	p, err := s.UpdateByID(id, ctx)

	if err != nil {
		ctx.AbortWithStatus(400) // Abort with :中止する
		fmt.Println(err)
	} else {
		ctx.JSON(200,p)
	}
}



// Delete action: DELETE /users/:id
func (pc Controller) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id") // ByNameメソッドは、与えられたKey(ここではid)からデータを取得する便利メソッド
	var s user.Service

	if err := s.DeleteByID(id); err != nil {
		ctx.AbortWithStatus(403) // Abort with :中止する
		fmt.Println(err)
	} else {
		ctx.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}