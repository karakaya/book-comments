package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func CategoryRoutes(router *mux.Router) {
	c := router.PathPrefix("/category").Subrouter()
	c.HandleFunc("/", controller.GetCategories).Methods("GET")
	c.HandleFunc("/", controller.CreateCategory).Methods("POST")
	c.HandleFunc("/{id}", controller.GetCategory).Methods("GET")
	c.HandleFunc("/{id}", controller.UpdateCategory).Methods("PATCH")
	c.HandleFunc("/{id}", controller.DeleteCategory).Methods("DELETE")
}
