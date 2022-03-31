package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/common/easy-lease/pkg/controllers"
	"github.com/gorilla/mux"
)

func TestCreateSociety(t *testing.T) {

	router := mux.NewRouter()

	router.HandleFunc("/societies", controllers.CreateSociety).Methods("POST")
	var jsonStr = []byte(`{   "societyAddress": "address1",
    "societyAmenities": "amenties1",
    "societyCity": "city1",
    "societyImg": "img1",
    "societyName": "name1"}`)
	request, _ := http.NewRequest("POST", "/societies", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	fmt.Println(response.Body.String())
	if m["SocietyAddress"] != "address1" {
		t.Errorf("Expected product address to be 'address1'. Got '%v'", m["SocietyAddress"])
	}

	if m["SocietyCity"] != "city1" {
		t.Errorf("Expected product city to be 'city1'. Got '%v'", m["SocietyCity"])
	}

}
