package routes

import (
	"github.com/gorilla/mux"
	//"Users/vamsiviswanathbethamsetty/Desktop/assignments/spring22/se/go-bookstore/pkg/controllers"
	"github.com/common/easy-lease/pkg/controllers"
)

var RegisterLeaseRoutes = func(router *mux.Router) {
	router.HandleFunc("/leases", controllers.CreateLease).Methods("POST")
	router.HandleFunc("/leases", controllers.GetLease).Methods("GET")
	router.HandleFunc("/leases/{leaseId}", controllers.GetLeaseById).Methods("GET")
	router.HandleFunc("/leases/{leaseId}", controllers.UpdateLease).Methods("PUT")
	router.HandleFunc("/leases/{leaseId}", controllers.DeleteLease).Methods("DELETE")
	//router.HandleFunc("/user/{email}", controllers.GetUserByEmail).Methods("GET")
}
