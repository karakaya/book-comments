package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func AskRoutes(r *mux.Router) {
	ar := r.PathPrefix("/ask").Subrouter()
	ar.HandleFunc("/", controller.GetAsks).Methods("GET")
	ar.HandleFunc("/{id}", controller.GetAsk).Methods("GET")
	ar.HandleFunc("/", controller.CreateAsk).Methods("POST")
	ar.HandleFunc("/{id}", controller.UpdateAsk).Methods("PATCH")
	ar.HandleFunc("/{id}", controller.DeleteAsk).Methods("DELETE")
}
