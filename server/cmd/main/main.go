package main

import (
	"log"
	"net/http"

	"github.com/common/easy-lease/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterLeaseRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
