package routes

import (
	"github.com/gorilla/mux"
	//"Users/vamsiviswanathbethamsetty/Desktop/assignments/spring22/se/go-bookstore/pkg/controllers"
	"github.com/common/easy-lease/pkg/controllers"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	//router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{email}", controllers.GetUserByEmail).Methods("GET")
}
