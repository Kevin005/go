package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"xiaoshi/util"
	"log"
	"net/http"
	"fmt"
	"xiaoshi/app/model"
)

type App struct {
	Router *mux.Router
	jump   *Jump
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
		return
	}
	a.Router = mux.NewRouter()
	a.jump = &Jump{db: model.DBMigrate(db)}
	a.setRouters()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) setRouters() {
	/** 意见反馈 */
	a.Post("/xiaoshi/feedback", a.jump.CreateFeedback)
	a.Get("/xiaoshi/feedback", a.jump.GetAllFeedback)
	/** 用户 */
	a.Post("/xiaoshi/avatar", a.jump.UploadAvatar)
	a.Get("/xiaoshi/avatar/{name}", a.jump.GetAvatar)
	a.Post("/xiaoshi/user/register", a.jump.Register)
	a.Post("/xiaoshi/user/login", a.jump.Login)
	a.Post("/xiaoshi/user/edit_pwd", a.jump.EditPwd)
	a.Post("/xiaoshi/user/edit", a.jump.EditUserInfo)
	/** 书籍 */
	a.Post("/xiaoshi/book", a.jump.CreateBook)
	a.Get("/xiaoshi/books/{user_id}", a.jump.getMyBooks)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
