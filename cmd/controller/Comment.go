package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//TODO : auth user and check the comment

func GetComment(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment
	err := database.DB.First(&comment, mux.Vars(r)["id"]).Error
	if err != nil {
		log.Printf("err get comment : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(comment)
	w.Write(res)
}
func GetComments(w http.ResponseWriter, r *http.Request) {
	var comments []model.Comment
	err := database.DB.Find(&comments).Error
	if err != nil {
		log.Printf("err get comments : %v ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(comments)
	w.Write(res)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	errFindBook := database.DB.First(&book, r.FormValue("book")).Error
	if errFindBook != nil {
		log.Printf("err find book: %v ", errFindBook)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//TODO: Get user id from authenticated user
	errCreateComment := database.DB.Create(&model.Comment{
		UserID:  1,
		Book:    book,
		Comment: r.FormValue("comment"),
	}).Error
	if errCreateComment != nil {
		log.Printf("err create comment : %v", errCreateComment)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment
	errFindComment := database.DB.First(&comment, mux.Vars(r)["id"]).Error
	if errFindComment != nil {
		log.Printf("find comment err in update controller: %v", errFindComment)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	errJsonDec := json.NewDecoder(r.Body).Decode(&comment)
	if errJsonDec != nil {
		log.Printf("err json dec : %v \n", errJsonDec)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//TODO: check the comment, it is really belongs the authenticated user
	database.DB.Save(&comment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment
	errFindComment := database.DB.First(&comment, mux.Vars(r)["id"]).Error
	if errFindComment != nil {
		log.Printf("find comment err: %v", errFindComment)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//TODO: check the comment, it is really belongs the authenticated user
	err := database.DB.Model(&model.Comment{}).Delete(&comment).Error
	if err != nil {
		log.Printf("err delete comment: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
