package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"xiaoshi/util"
	"log"
	"net/http"
	"fmt"
	"xiaoshi/app/model"
	"xiaoshi/app/handler"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *util.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.DBName,
		config.DB.Charset)
	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRouters() {
	a.Post("/xiaoshi/feedback", a.CreateFeedback)
	a.Get("/xiaoshi/feedback", a.GetAllFeedback)
	a.Post("/xiaoshi/avatar", a.UploadAvatar)
	a.Get("/xiaoshi/avatar/{name}", a.GetAvatar)
	a.Post("/xiaoshi/user/register", a.Register)
	a.Post("/xiaoshi/user/login", a.Login)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

/**
反馈
 */
func (a *App) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	handler.CreateFeedback(a.DB, w, r)
}

/**
查询所有反馈
 */
func (a *App) GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	handler.GetAllFeedback(a.DB, w, r)
}

/**
上传用户头像
 */
func (a *App) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	handler.UploadAvatar(w, r)
}

/**
获取用户头像
 */
func (a *App) GetAvatar(w http.ResponseWriter, r *http.Request) {
	handler.GetAvatar(w, r)
}

/**
用户注册
 */
func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	handler.Register(a.DB, w, r)
}

/**
用户登录
 */
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	handler.Login(a.DB, w, r)
}
