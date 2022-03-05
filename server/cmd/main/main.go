package main

import (
	"log"
	"net/http"
	"os"

	"github.com/common/easy-lease/pkg/routes"
	"github.com/gorilla/handlers"
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

	corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	headersOk := handlers.AllowedHeaders([]string{"accept", "origin", "X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"})
	allowCreds := handlers.AllowCredentials()
	allowOptions := handlers.OptionStatusCode(204)

	r.HandleFunc("/", RootEndPoint).Methods("GET")

	port := os.Getenv("PORT")
	routes.RegisterUserRoutes(r)
	routes.RegisterLeaseRoutes(r)
	routes.RegisterApartmentRoutes(r)
	routes.RegisterListingRoutes(r)
	routes.LoginRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj, headersOk, methodsOk, allowCreds, allowOptions)(r)))

}
