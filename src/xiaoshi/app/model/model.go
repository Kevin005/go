package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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

type Users struct {
	Model
	CreateTime  time.Time `json:"create_time"`
	UserName    string    `json:"user_name"`
	NickName    string    `json:"nick_name"`
	PhoneNumber string    `json:"phone_number"`
	Age         int       `json:"age"`
	Gender      int       `json:"gender"`
	Token       string    `json:"token"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Feedbacks{}, &Users{})
	return db
}
