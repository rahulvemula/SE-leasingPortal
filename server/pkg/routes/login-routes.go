package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var LoginRoutes = func(router *mux.Router) {
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
}
