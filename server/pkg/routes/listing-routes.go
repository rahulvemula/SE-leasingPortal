package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterListingRoutes = func(router *mux.Router) {
	router.HandleFunc("/listing", controllers.CreateListing).Methods("POST")
	router.HandleFunc("/listing", controllers.GetListing).Methods("GET")
	router.HandleFunc("/listing/{listingId}", controllers.GetListingById).Methods("GET")
	router.HandleFunc("/listing/{listingId}", controllers.UpdateListing).Methods("PUT")
	router.HandleFunc("/listing/{listingId}", controllers.DeleteListing).Methods("DELETE")
}
