package main

import (
	"log"
	"net/http"
	"os"

	"github.com/common/easy-lease/pkg/routes"
	"github.com/gorilla/mux"
)

func Add(a int, b int) int {
	return a + b
}

func RootEndPoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", RootEndPoint).Methods("GET")

	port := os.Getenv("PORT")
	routes.RegisterUserRoutes(r)
	routes.RegisterLeaseRoutes(r)
	routes.RegisterApartmentRoutes(r)
	routes.RegisterListingRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
