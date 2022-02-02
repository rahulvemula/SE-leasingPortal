package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterApartmentRoutes = func(router *mux.Router) {
	router.HandleFunc("/apartment/", controllers.CreateApartment).Methods("POST")
	router.HandleFunc("/apartment/", controllers.GetApartment).Methods("GET")
	router.HandleFunc("/apartment/{apartmentId}", controllers.GetApartmentById).Methods("GET")
	router.HandleFunc("/apartment/{apartmentId}", controllers.UpdateApartment).Methods("PUT")
	router.HandleFunc("/apartment/{apartmentId}", controllers.DeleteApartment).Methods("DELETE")
}
