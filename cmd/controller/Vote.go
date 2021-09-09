package controller

import (
	"book-comments/cmd/database"
	"book-comments/cmd/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//TODO: Add user check to prevent multi vote
func Upvote(w http.ResponseWriter, r *http.Request) {
	voteType := mux.Vars(r)["type"]
	id := mux.Vars(r)["id"]
	if voteType == "book" {
		var book model.Book
		database.DB.First(&book, id)
		book.VoteCount++
		database.DB.Save(&book)
	} else if voteType == "ask" {
		var ask model.Ask
		database.DB.First(&ask, id)
		ask.VoteCount++
		database.DB.Save(&ask)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("there is no %s to vote", voteType)
		return
	}

}
