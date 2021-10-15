package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAsks(w http.ResponseWriter, r *http.Request) {
	var asks []model.Ask
	err := database.DB.Find(&asks).Error
	if err != nil {
		log.Printf("err get ask's: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(asks)
	w.Write(res)
}

func GetAsk(w http.ResponseWriter, r *http.Request) {
	var ask model.Ask
	err := database.DB.First(&ask, mux.Vars(r)["id"]).Error
	if err != nil {
		log.Printf("err get ask: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(ask)
	w.Write(res)
}

func CreateAsk(w http.ResponseWriter, r *http.Request) {
	//TODO: Get user id from authenticated user
	err := database.DB.Create(&model.Ask{
		Title:   r.FormValue("title"),
		UserID:  1,
		Content: r.FormValue("content"),
	}).Error

	if err != nil {
		log.Printf("err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateAsk(w http.ResponseWriter, r *http.Request) {
	var ask model.Ask
	err := database.DB.First(&ask, mux.Vars(r)["id"]).Error
	if err != nil {
		log.Printf("err update ask: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//TODO: check the ask is it belongs to authenticated user
	errJson := json.NewDecoder(r.Body).Decode(&ask)
	if errJson != nil {
		log.Printf("err: %v \n", errJson)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func DeleteAsk(w http.ResponseWriter, r *http.Request) {
	var ask model.Ask
	err := database.DB.First(&ask, mux.Vars(r)["id"]).Error
	if err != nil {
		log.Printf("err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//TODO: check the ask is it belongs to authenticated user
	errDelete := database.DB.Model(&model.Ask{}).Delete(&ask).Error
	if errDelete != nil {
		log.Printf("err delete :%v", errDelete)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
