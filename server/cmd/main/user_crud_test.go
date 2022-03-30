package main

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

func TestCreateUser(t *testing.T) {

	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	var jsonStr = []byte(`{"email":"sai@ufl.edu", "name": "sai teja", "password":"sai"}`)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["Name"] != "sai teja" {
		t.Errorf("Expected product name to be 'sai teja'. Got '%v'", m["Name"])
	}

	if m["Email"] != "sai@ufl.edu" {
		t.Errorf("Expected product email to be 'sai@ufl.edu'. Got '%v'", m["Email"])
	}

}
