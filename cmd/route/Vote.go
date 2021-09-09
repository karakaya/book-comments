package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func VoteRoutes(r *mux.Router) {
	v := r.PathPrefix("/vote").Subrouter()
	v.HandleFunc("/{type}/{id}", controller.Upvote).Methods("GET")
}
