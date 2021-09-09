package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	err := database.DB.Preload("Category").Preload("Author").Find(&books).Error
	if err != nil {
		log.Printf("get book err %v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, jsonErr := json.Marshal(books)
	if jsonErr != nil {
		log.Printf("book json marshal err %v \n", jsonErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var book model.Book
	err := database.DB.Preload("Category").Preload("Author").First(&book, id).Error
	if err != nil {
		log.Printf("err: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, jsonErr := json.Marshal(book)
	if jsonErr != nil {
		log.Printf("err %v \n", jsonErr)
		return
	}
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	errJsonDec := json.NewDecoder(r.Body).Decode(&book)
	if errJsonDec != nil {
		log.Printf("create book json dec err : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Create(&book).Error
	if err != nil {
		log.Printf("err : %v \n", err)
		return
	}

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var book model.Book
	database.DB.First(&book, id)
	errJsonDec := json.NewDecoder(r.Body).Decode(&book)
	if errJsonDec != nil {
		log.Printf("err update json dec : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err := database.DB.Save(&book).Error
	if err != nil {
		log.Printf("err book update : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(book)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := database.DB.Find(&model.Book{}, id).Delete(&id).Error
	if err != nil {
		log.Printf("delete book err : %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
