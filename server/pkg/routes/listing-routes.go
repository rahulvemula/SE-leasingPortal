package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterListingRoutes = func(router *mux.Router) {
	router.HandleFunc("/listings", controllers.CreateListing).Methods("POST")
	router.HandleFunc("/listings", controllers.GetListing).Methods("GET")
	router.HandleFunc("/listings/{listingId}", controllers.GetListingById).Methods("GET")
	router.HandleFunc("/listings/{listingId}", controllers.UpdateListing).Methods("PUT")
	router.HandleFunc("/listings/{listingId}", controllers.DeleteListing).Methods("DELETE")
}
