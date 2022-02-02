package routes

import (
	"github.com/gorilla/mux"
	//"Users/vamsiviswanathbethamsetty/Desktop/assignments/spring22/se/go-bookstore/pkg/controllers"
	"github.com/common/easy-lease/pkg/controllers"
)

var RegisterLeaseRoutes = func(router *mux.Router) {
	router.HandleFunc("/lease/", controllers.CreateLease).Methods("POST")
	router.HandleFunc("/lease/", controllers.GetLease).Methods("GET")
	router.HandleFunc("/lease/{leaseId}", controllers.GetLeaseById).Methods("GET")
	router.HandleFunc("/lease/{leaseId}", controllers.UpdateLease).Methods("PUT")
	router.HandleFunc("/lease/{leaseId}", controllers.DeleteLease).Methods("DELETE")
	//router.HandleFunc("/user/{email}", controllers.GetUserByEmail).Methods("GET")
}
