// User model(entity/user.go内)の動きを定義
// GetAll, CreateModel, GetByID, UpdateByID, DeleteByID
package user

import (
	"github.com/gin-gonic/gin"

	"go-backend-tuto/db"
	"go-backend-tuto/entity"
)

// Service procides user's behavior
// ユーザーの行動を把握するサービス
type Service struct{}

// User is alias of entity.User struct
// User は、entity.User 構造体のエイリアスです。
type User entity.User

// GetAll is get all User
func (s Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// CreateModel is create User model
func (s Service) CreateModel(ctx *gin.Context) (User, error){
	db := db.GetDB()
	var u User

	if err := ctx.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// GetByID is get a User
func (s Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	
	return u, nil
}

func (s Service) UpdateByID(id string, ctx *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	
	if err := ctx.BindJSON(&u); err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}
	
	return nil
}