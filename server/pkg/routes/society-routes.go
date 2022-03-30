package routes

import (
	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterSocietyRoutes = func(router *mux.Router) {
	router.Handle("/societies", controllers.IsAuthorized(controllers.CreateSociety)).Methods("POST")
	router.Handle("/societies", controllers.IsAuthorized(controllers.GetSociety)).Methods("GET")
	router.Handle("/societies/{societyId}", controllers.IsAuthorized(controllers.GetSocietyById)).Methods("GET")
	router.Handle("/societies/{societyId}", controllers.IsAuthorized(controllers.UpdateSociety)).Methods("PUT")
	router.Handle("/societies/{societyId}", controllers.IsAuthorized(controllers.DeleteSociety)).Methods("DELETE")
}
