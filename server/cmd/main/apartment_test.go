package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/common/easy-lease/pkg/controllers"
	"github.com/common/easy-lease/pkg/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetApartments(t *testing.T) {
	router := mux.NewRouter()
	var apartments []models.Apartment

	router.HandleFunc("/apartment/", controllers.GetApartment).Methods("GET")
	request, _ := http.NewRequest("GET", "/apartment/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &apartments)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(apartments) >= 1 {
		name := apartments[1].Name
		address := apartments[1].Address
		//fmt.Println(startDateNotNull, endDateNotNull)
		assert.Equal(t, "Stoneridge", name, "Non empty name expected")
		assert.Equal(t, "3800 SW 34th St", address, "Non empty address expected")
	}
}

func TestGetApartmentById(t *testing.T) {
	router := mux.NewRouter()
	var apartment models.Apartment

	router.HandleFunc("/apartment/{apartmentId}", controllers.GetApartmentById).Methods("GET")
	request, _ := http.NewRequest("GET", "/apartment/2", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &apartment)

	fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}

	if (apartment != models.Apartment{}) {
		assert.Equal(t, 2, int(apartment.ID), "ID's should be matched")
	}
}
