package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterListingRoutes = func(router *mux.Router) {
	router.Handle("/listings", controllers.IsAuthorized(controllers.CreateListing)).Methods("POST")
	router.Handle("/listings", controllers.IsAuthorized(controllers.GetListing)).Methods("GET")
	router.Handle("/listings/{listingId}", controllers.IsAuthorized(controllers.GetListingById)).Methods("GET")
	router.Handle("/listings/{listingId}", controllers.IsAuthorized(controllers.UpdateListing)).Methods("PUT")
	router.Handle("/listings/{listingId}", controllers.IsAuthorized(controllers.DeleteListing)).Methods("DELETE")
}
