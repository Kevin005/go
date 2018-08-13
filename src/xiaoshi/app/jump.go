package app

import (
	"net/http"
	"xiaoshi/app/handler"
	"github.com/jinzhu/gorm"
)

type Jump struct {
	db *gorm.DB
}

/**
反馈
 */
func (j *Jump) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	handler.CreateFeedback(j.db, w, r)
}

/**
查询所有反馈
 */
func (j *Jump) GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	handler.GetAllFeedback(j.db, w, r)
}

/**
上传用户头像
 */
func (j *Jump) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	handler.UploadAvatar(w, r)
}

/**
获取用户头像
 */
func (j *Jump) GetAvatar(w http.ResponseWriter, r *http.Request) {
	handler.GetAvatar(w, r)
}

/**
用户注册
 */
func (j *Jump) Register(w http.ResponseWriter, r *http.Request) {
	handler.Register(j.db, w, r)
}

/**
用户登录
 */
func (j *Jump) Login(w http.ResponseWriter, r *http.Request) {
	handler.Login(j.db, w, r)
}

/**
修改密码
 */
func (j *Jump) EditPwd(w http.ResponseWriter, r *http.Request) {
	handler.EditPwd(j.db, w, r)
}

/**
修改用户信息
 */
func (j *Jump) EditUserInfo(w http.ResponseWriter, r *http.Request) {
	handler.EditUserInfo(j.db, w, r)
}

/**
发布一本书
 */
func (j *Jump) CreateBook(w http.ResponseWriter, r *http.Request) {
	handler.CreateBook(j.db, w, r)
}

/**
获取我所有在读的书
 */
func (j *Jump) getMyBooks(w http.ResponseWriter, r *http.Request) {
	handler.GetMyBooks(j.db, w, r)
}
