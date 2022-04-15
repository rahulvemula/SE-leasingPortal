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

func TestGetSocieties(t *testing.T) {
	router := mux.NewRouter()
	var societies []models.Society

	router.HandleFunc("/societies", controllers.GetSociety).Methods("GET")
	request, _ := http.NewRequest("GET", "/societies", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &societies)

	if err != nil {
		fmt.Println("err is ", err)
	}
	//fmt.Println(societies[1].SocietyName)
	if len(societies) >= 1 {
		name := societies[0].SocietyName
		address := societies[0].SocietyAddress

		//fmt.Println(startDateNotNull, endDateNotNull)
		assert.Equal(t, len(name) > 0, true, "Non empty name expected")
		assert.Equal(t, len(address) > 0, true, "Non empty address expected")
	}
}

func TestGetSocietyById(t *testing.T) {
	router := mux.NewRouter()
	var society models.Society

	router.HandleFunc("/societies/{societyId}", controllers.GetSocietyById).Methods("GET")
	request, _ := http.NewRequest("GET", "/societies/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &society)

	fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}

	if (society != models.Society{}) {
		assert.Equal(t, 1, int(society.ID), "ID's should be matched")
	}
}
