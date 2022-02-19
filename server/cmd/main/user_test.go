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

func TestGetUsers(t *testing.T) {
	router := mux.NewRouter()
	var users []models.User

	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	request, _ := http.NewRequest("GET", "/user/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &users)

	if err != nil {
		fmt.Println("err is ", err)
	}

	if len(users) >= 1 {
		nameNotNull := len(users[0].Name) > 0
		assert.Equal(t, true, nameNotNull, "Non empty name expected")
	}
}

func TestGetUserByEmail(t *testing.T) {
	router := mux.NewRouter()
	var user models.User

	router.HandleFunc("/user/{email}", controllers.GetUserByEmail).Methods("GET")
	request, _ := http.NewRequest("GET", "/user/vbethamsetty@ufl.edu", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	//testing for just status code here
	assert.Equal(t, 200, response.Code, "OK response is expected")
	err := json.Unmarshal([]byte(response.Body.Bytes()), &user)

	//fmt.Println("response is: ", response.Body.String())
	if err != nil {
		fmt.Println("err is ", err)
	}

	if (user != models.User{}) {
		//fmt.Println(user.Email)
		assert.Equal(t, "vbethamsetty@ufl.edu", user.Email, "Emails should be matching")
	}
}
