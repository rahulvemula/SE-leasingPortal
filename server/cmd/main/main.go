package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/common/easy-lease/cmd/main/docs"

	httpSwagger "github.com/swaggo/http-swagger"

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

// @title LeasingPortal APIs
// @version 1.0
// @description This is a sample service for managing leasing apis
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	headersOk := handlers.AllowedHeaders([]string{"accept", "origin", "X-Requested-With", "Content-Type", "Authorization", "token"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"})
	allowCreds := handlers.AllowCredentials()
	allowOptions := handlers.OptionStatusCode(204)

	r.HandleFunc("/", RootEndPoint).Methods("GET")

	port := os.Getenv("PORT")
	routes.RegisterUserRoutes(r)
	routes.RegisterLeaseRoutes(r)
	routes.RegisterApartmentRoutes(r)
	routes.RegisterListingRoutes(r)
	routes.RegisterComplaintRoutes(r)
	routes.RegisterSocietyRoutes(r)
	routes.LoginRoutes(r)

	http.Handle("/", r)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj, headersOk, methodsOk, allowCreds, allowOptions)(r)))
	// log.Fatal(http.ListenAndServe("localhost:9011", handlers.CORS(corsObj, headersOk, methodsOk, allowCreds, allowOptions)(r)))

}
