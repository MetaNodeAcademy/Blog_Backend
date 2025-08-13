package repo

import (
	"Blog_Backend/config"
	"Blog_Backend/model"
	"context"
	"gorm.io/gorm"
)

func CreateUser(r *model.User) {
	config.GetDBConn().Create(&r)
}

func FindUser(email string) string {
	first, err := gorm.G[model.User](config.GetDBConn()).Where("email = ?", email).First(context.Background())
	if err != nil {
		return ""
	} else {
		return first.Password
	}

}
