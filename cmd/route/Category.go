package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func CategoryRoutes(router *mux.Router) {
	router.PathPrefix("/category").Subrouter()
	router.HandleFunc("/", controller.GetCategories).Methods("GET")
	router.HandleFunc("/", controller.CreateCategory).Methods("POST")
	router.HandleFunc("/{id}", controller.GetCategory).Methods("GET")
	router.HandleFunc("/{id}", controller.UpdateCategory).Methods("PATCH")
	router.HandleFunc("/{id}", controller.DeleteCategory).Methods("DELETE")
}
