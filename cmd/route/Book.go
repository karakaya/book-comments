package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	b := router.PathPrefix("/book").Subrouter()
	b.HandleFunc("/", controller.GetBooks).Methods("GET")
	b.HandleFunc("/", controller.CreateBook).Methods("POST")
	b.HandleFunc("/{id}", controller.GetBook).Methods("GET")
	b.HandleFunc("/{id}", controller.UpdateBook).Methods("PATCH")
	b.HandleFunc("/{id}", controller.DeleteBook).Methods("DELETE")
}
