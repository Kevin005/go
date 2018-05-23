package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}

type Feedbacks struct {
	Model
	UserId  int    `json:"user_id"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Feedbacks{})
	return db
}
