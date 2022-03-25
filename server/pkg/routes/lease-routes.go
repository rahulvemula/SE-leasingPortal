package routes

import (
	"github.com/gorilla/mux"
	//"Users/vamsiviswanathbethamsetty/Desktop/assignments/spring22/se/go-bookstore/pkg/controllers"
	"github.com/common/easy-lease/pkg/controllers"
)

var RegisterLeaseRoutes = func(router *mux.Router) {
	router.Handle("/leases", controllers.IsAuthorized(controllers.CreateLease)).Methods("POST")
	router.Handle("/leases", controllers.IsAuthorized(controllers.GetLease)).Methods("GET")
	router.Handle("/leases/{leaseId}", controllers.IsAuthorized(controllers.GetLeaseById)).Methods("GET")
	router.Handle("/leases/{leaseId}", controllers.IsAuthorized(controllers.UpdateLease)).Methods("PUT")
	router.Handle("/leases/{leaseId}", controllers.IsAuthorized(controllers.DeleteLease)).Methods("DELETE")
	//router.HandleFunc("/user/{email}", controllers.GetUserByEmail).Methods("GET")
}
