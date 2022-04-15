package main

import (
	"bytes"
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

func TestGetListings(t *testing.T) {
	router := mux.NewRouter()
	var listings []models.Listing

	router.HandleFunc("/listing/", controllers.GetListing).Methods("GET")
	request, _ := http.NewRequest("GET", "/listing/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &listings)

	if err != nil {
		fmt.Println("err is ", err)
	}

	fmt.Println(listings)

	if len(listings) >= 1 {
		rent := listings[0].Rent
		houseType := listings[0].HouseType
		assert.Equal(t, int64(500), rent, "Non empty name expected")
		assert.Equal(t, "1", houseType, "Non empty address expected")
	}
}

func TestGetListingById(t *testing.T) {
	router := mux.NewRouter()
	var listing models.Listing

	router.HandleFunc("/listing/{listingId}", controllers.GetListingById).Methods("GET")
	request, _ := http.NewRequest("GET", "/listing/3", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &listing)

	fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}

	if (listing != models.Listing{}) {
		assert.Equal(t, 3, int(listing.ID), "ID's should be matched")
	}
}

func TestCreateListing(t *testing.T) {

	router := mux.NewRouter()

	router.HandleFunc("/listings", controllers.CreateListing).Methods("POST")
	var jsonStr = []byte(`{"houseType": "2bhk",
	"isleased": true,
	"listingImg": "randomimg",
	"listingType": "randomtype",
	"rent": 1000,
	"societyId": "randomid",
	"userId": 11}`)
	request, _ := http.NewRequest("POST", "/listings", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	fmt.Println(response.Body.String())

	if m["houseType"] != "2bhk" {
		t.Errorf("Expected product city to be '2bhk'. Got '%v'", m["houseType"])
	}

}
