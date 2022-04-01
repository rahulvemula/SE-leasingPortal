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

func TestGetComplaints(t *testing.T) {
	router := mux.NewRouter()
	var complaints []models.Complaint

	router.HandleFunc("/complaint", controllers.GetComplaints).Methods("GET")
	request, _ := http.NewRequest("GET", "/complaint", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &complaints)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(complaints) >= 1 {
		description := complaints[0].Description
		assert.Equal(t, "First Complaint", description, "Non empty name expected")
	}
}

func TestGetComplaintByUser(t *testing.T) {
	router := mux.NewRouter()
	var complaints []models.Complaint

	router.HandleFunc("/complaints/user/{userId}", controllers.GetComplaintsByUser).Methods("GET")
	request, _ := http.NewRequest("GET", "/complaints/user/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &complaints)

	fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(complaints) >= 1 {
		description := complaints[0].Description
		assert.Equal(t, "First Complaint", description, "Non empty name expected")
	}
}
