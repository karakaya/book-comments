package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.PathPrefix("/book").Subrouter()
	router.HandleFunc("/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/{id}", controller.UpdateBook).Methods("PATCH")
	router.HandleFunc("/{id}", controller.DeleteBook).Methods("DELETE")
}
