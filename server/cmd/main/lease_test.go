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

func TestGetLeases(t *testing.T) {
	router := mux.NewRouter()
	var leases []models.Lease

	router.HandleFunc("/lease/", controllers.GetLease).Methods("GET")
	request, _ := http.NewRequest("GET", "/lease/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &leases)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(leases) >= 1 {
		startDateNotNull := len(leases[0].LeaseStartDate) > 0
		endDateNotNull := len(leases[0].LeaseEndDate) > 0
		//fmt.Println(startDateNotNull, endDateNotNull)
		assert.Equal(t, true, startDateNotNull, "Non empty name expected")
		assert.Equal(t, true, endDateNotNull, "Non empty name expected")
	}
}

func TestGetLeaseById(t *testing.T) {
	router := mux.NewRouter()
	var lease models.Lease

	router.HandleFunc("/lease/{leaseId}", controllers.GetLeaseById).Methods("GET")
	request, _ := http.NewRequest("GET", "/lease/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &lease)

	fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}

	if (lease != models.Lease{}) {
		assert.Equal(t, 1, int(lease.ID), "ID's should be matched")
	}
}

func TestCreateLease(t *testing.T) {

	router := mux.NewRouter()

	router.HandleFunc("/leases", controllers.CreateLease).Methods("POST")
	var jsonStr = []byte(`{  "leaseEndDate": "20aug2022",
    "leaseStartDate": "20aug2019",
    "listingId": 11,
    "userId": 33}`)
	request, _ := http.NewRequest("POST", "/leases", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	fmt.Println(response.Body.String())

	if m["LeaseStartDate"] != "20aug2019" {
		t.Errorf("Expected lease start date to be '20aug2019'. Got '%v'", m["LeaseStartDate"])
	}

}
