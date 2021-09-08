package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func CommentRoutes(r *mux.Router) {
	cm := r.PathPrefix("/comment").Subrouter()
	cm.HandleFunc("/", controller.GetComments).Methods("GET")
	cm.HandleFunc("/", controller.CreateComment).Methods("POST")
	cm.HandleFunc("/{id}", controller.GetComment).Methods("GET")
	cm.HandleFunc("/{id}", controller.UpdateComment).Methods("PATCH")
	cm.HandleFunc("/{id}", controller.DeleteComment).Methods("DELETE")
}
