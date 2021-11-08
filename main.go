package main

import (
	"book-comments/cmd/database"
	"book-comments/cmd/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	route.AuthorRoutes(router)
	route.BookRoutes(router)
	route.CategoryRoutes(router)
	route.CommentRoutes(router)
	route.AskRoutes(router)
	route.VoteRoutes(router)
	route.UserRoute(router)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	database.ConnectDB()
	database.Migrate()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("http listen err: %v", err)
	}

}
