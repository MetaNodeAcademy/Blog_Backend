package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title     string
	Content   string
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	gorm.Model
	Content   string
	UserId    uint
	ArticleId uint
	CreatedAt time.Time
}

type User struct {
	gorm.Model
	UserName string
	Password string
	Email    string
	Articles []Article
}
