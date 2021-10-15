package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categorys []model.Category
	err := database.DB.Find(&categorys).Error
	if err != nil {
		log.Printf("get category err %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, jsonErr := json.Marshal(categorys)
	if jsonErr != nil {
		log.Printf("categories json marshal err %v \n", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var category model.Category
	err := database.DB.First(&category, id).Error
	if err != nil {
		log.Printf("err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, jsonErr := json.Marshal(category)
	if jsonErr != nil {
		log.Printf("err %v \n", jsonErr)
		return
	}
	w.Write(res)

}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	errJsonDec := json.NewDecoder(r.Body).Decode(&category)
	if errJsonDec != nil {
		log.Printf("create category json dec err : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Create(&category).Error
	if err != nil {
		log.Printf("err : %v \n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var category model.Category
	database.DB.First(&category, id)
	errJsonDec := json.NewDecoder(r.Body).Decode(&category)
	if errJsonDec != nil {
		log.Printf("err update json dec : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Save(&category).Error
	if err != nil {
		log.Printf("err category update : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(category)
	w.Write(res)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := database.DB.Find(&model.Category{}, id).Delete(&id).Error
	if err != nil {
		log.Printf("delete category err : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
