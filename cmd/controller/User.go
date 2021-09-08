package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Printf("get user err %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, jsonErr := json.Marshal(users)
	if jsonErr != nil {
		log.Printf("user json marshal err %v \n", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user model.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		log.Printf("err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		log.Printf("err %v \n", jsonErr)
		return
	}
	w.Write(res)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	errJsonDec := json.NewDecoder(r.Body).Decode(&user)
	if errJsonDec != nil {
		log.Printf("create user json dec err : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Create(&user).Error
	if err != nil {
		log.Printf("err : %v \n", err)
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user model.User
	database.DB.First(&user, id)
	errJsonDec := json.NewDecoder(r.Body).Decode(&user)
	if errJsonDec != nil {
		log.Printf("err update json dec : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Save(&user).Error
	if err != nil {
		log.Printf("err user update : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(user)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := database.DB.Find(&model.User{}, id).Delete(&id).Error
	if err != nil {
		log.Printf("delete user err : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
