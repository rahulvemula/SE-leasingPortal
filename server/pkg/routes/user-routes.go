package routes

import (
	"github.com/gorilla/mux"
	//"Users/vamsiviswanathbethamsetty/Desktop/assignments/spring22/se/go-bookstore/pkg/controllers"
	"github.com/common/easy-lease/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.Handle("/users", controllers.IsAuthorized(controllers.GetUser)).Methods("GET")
	//router.Handle("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.Handle("/users/{userId}", controllers.IsAuthorized(controllers.UpdateUser)).Methods("PUT")
	router.Handle("/users/{userId}", controllers.IsAuthorized(controllers.DeleteUser)).Methods("DELETE")
	router.Handle("/users/{email}", controllers.IsAuthorized(controllers.GetUserByEmail)).Methods("GET")
}
