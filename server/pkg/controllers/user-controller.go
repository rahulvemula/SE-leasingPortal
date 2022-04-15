package controllers

import (
	//"encoded/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/common/easy-lease/pkg/models"
	"github.com/common/easy-lease/pkg/utils"
	"github.com/gorilla/mux"
)

var NewUser models.User

// GetUsers godoc
// @Summary Get details of all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
//@Param token header string true "token header"
// @Success 200 {array} models.User
// @Router /users [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, _ := models.GetUserById(ID)

	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// GetUser godoc
// @Summary Get details of the user by email
// @Description Get details of user by email
// @Tags users
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param email path string true "email"
// @Success 200 {array} models.User
// @Router /users/{email} [get]
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	userDetails, _ := models.GetUserByEmail(email)

	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Create godoc
// @Summary Create a new
// @Description Get a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param body body string true "body"
// @Success 200 {array} models.User
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	u := CreateUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Delete user godoc
// @Summary Delete user
// @Description Delete the existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param userId path string true "userId"
// @Success 200 {array} models.User
// @Router /users/{userId} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Create godoc
// @Summary Update the new
// @Description Update the user
// @Tags users
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param userId path string true "userId"
// @Param body body string true "body"
// @Success 200 {array} models.User
// @Router /users/{userId} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	userDetails, db := models.GetUserById(ID)

	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}

	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		userDetails.Password = updateUser.Password
	}

	db.Save(&userDetails)

	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
