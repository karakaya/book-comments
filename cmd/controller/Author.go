package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//TODO : Add user check after adding jwt, for to be ensure user.

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author
	err := database.DB.Find(&authors).Error
	if err != nil {
		log.Printf("get author err %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, jsonErr := json.Marshal(authors)
	if jsonErr != nil {
		log.Printf("author json marshal err %v \n", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var author model.Author
	err := database.DB.First(&author, id).Error
	if err != nil {
		log.Printf("err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, jsonErr := json.Marshal(author)
	if jsonErr != nil {
		log.Printf("err %v \n", jsonErr)
		return
	}
	w.Write(res)

}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	errJsonDec := json.NewDecoder(r.Body).Decode(&author)
	if errJsonDec != nil {
		log.Printf("create author json dec err : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Create(&author).Error
	if err != nil {
		log.Printf("err : %v \n", err)
		return
	}

}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var author model.Author
	database.DB.First(&author, id)
	errJsonDec := json.NewDecoder(r.Body).Decode(&author)
	if errJsonDec != nil {
		log.Printf("err update json dec : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Save(&author).Error
	if err != nil {
		log.Printf("err author update : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(author)
	w.Write(res)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := database.DB.Find(&model.Author{}, id).Delete(&id).Error
	if err != nil {
		log.Printf("delete author err : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
