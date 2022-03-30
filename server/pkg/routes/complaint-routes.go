package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterComplaintRoutes = func(router *mux.Router) {
	router.Handle("/complaints", controllers.IsAuthorized(controllers.CreateComplaint)).Methods("POST")
	router.Handle("/complaints", controllers.IsAuthorized(controllers.GetComplaints)).Methods("GET")
	router.Handle("/complaints/{complaintId}", controllers.IsAuthorized(controllers.GetComplaintById)).Methods("GET")
	router.Handle("/complaints/user/{userId}", controllers.IsAuthorized(controllers.GetComplaintsByUser)).Methods("GET")
	router.Handle("/complaints/{complaintId}", controllers.IsAuthorized(controllers.DeleteComplaint)).Methods("DELETE")
}
