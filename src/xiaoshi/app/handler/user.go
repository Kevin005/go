package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"xiaoshi/app/model/response"
	"xiaoshi/app/model"
	"encoding/json"
)

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respUser := response.RespUser{}
	respUser.Data = user
	respUser.Message = "pass"
	respUser.Success = "0"
	respondJSON(w, http.StatusCreated, respUser)
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	respUser := response.RespUser{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}
	defer r.Body.Close()
	if checkToken(db, user.Token) {
		db.First(&user, "token = ?", user.Token)
		respUserData := response.RespUserData{}
		respUserData.User = user
		respUser.Data = respUserData
		respUser.Message = "pass"
		respUser.Success = "0"
		respondJSON(w, http.StatusCreated, respUser)
	} else {
		respUser := response.RespUser{}
		respUser.Message = "reject"
		respUser.Success = "1"
		respUser.Data = "use not found"
		respondJSON(w, http.StatusGone, respUser)
	}

}
