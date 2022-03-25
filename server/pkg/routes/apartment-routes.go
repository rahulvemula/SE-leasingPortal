package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterApartmentRoutes = func(router *mux.Router) {
	router.Handle("/apartments", controllers.IsAuthorized(controllers.CreateApartment)).Methods("POST")
	router.Handle("/apartments", controllers.IsAuthorized(controllers.GetApartment)).Methods("GET")
	router.Handle("/apartments/{apartmentId}", controllers.IsAuthorized(controllers.GetApartmentById)).Methods("GET")
	router.Handle("/apartments/{apartmentId}", controllers.IsAuthorized(controllers.UpdateApartment)).Methods("PUT")
	router.Handle("/apartments/{apartmentId}", controllers.IsAuthorized(controllers.DeleteApartment)).Methods("DELETE")
}
