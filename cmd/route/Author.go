package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(router *mux.Router) {
	router.PathPrefix("/author").Subrouter()
	router.HandleFunc("/", controller.GetAuthors).Methods("GET")
	router.HandleFunc("/", controller.CreateAuthor).Methods("POST")
	router.HandleFunc("/{id}", controller.GetAuthor).Methods("GET")
	router.HandleFunc("/{id}", controller.UpdateAuthor).Methods("PATCH")
	router.HandleFunc("/{id}", controller.DeleteAuthor).Methods("DELETE")
}
