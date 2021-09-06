package main

import (
	"book-comments/cmd/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	route.BookRoutes(router)
	route.AuthorRoutes(router)
	route.CategoryRoutes(router)
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatalf("http listen err: %v", err)
	}

}
