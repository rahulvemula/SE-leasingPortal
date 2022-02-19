package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.Equal(t, 4, total, "Expecting 4")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndPoint).Methods("GET")
	return router
}

func TestRootEndPoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	//fmt.Println(response.Body)
	//assert.Equal(t, "Hello World", response.Body.String(), "Hello world message is expected")
}
