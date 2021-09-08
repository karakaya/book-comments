package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(router *mux.Router) {
	a := router.PathPrefix("/author").Subrouter().StrictSlash(true)
	a.HandleFunc("/", controller.GetAuthors).Methods("GET")
	a.HandleFunc("/", controller.CreateAuthor).Methods("POST")
	a.HandleFunc("/{id}", controller.GetAuthor).Methods("GET")
	a.HandleFunc("/{id}", controller.UpdateAuthor).Methods("PATCH")
	a.HandleFunc("/{id}", controller.DeleteAuthor).Methods("DELETE")

}
