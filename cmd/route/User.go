package route

import (
	"book-comments/cmd/controller"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	u := r.PathPrefix("/user").Subrouter().StrictSlash(true)
	u.HandleFunc("/register", controller.CreateUser)
}
