package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterApartmentRoutes = func(router *mux.Router) {
	router.HandleFunc("/apartments", controllers.CreateApartment).Methods("POST")
	router.HandleFunc("/apartments", controllers.GetApartment).Methods("GET")
	router.HandleFunc("/apartments/{apartmentId}", controllers.GetApartmentById).Methods("GET")
	router.HandleFunc("/apartments/{apartmentId}", controllers.UpdateApartment).Methods("PUT")
	router.HandleFunc("/apartments/{apartmentId}", controllers.DeleteApartment).Methods("DELETE")
}
