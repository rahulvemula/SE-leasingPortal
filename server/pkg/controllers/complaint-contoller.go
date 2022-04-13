package controllers

import (
	//"encoded/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/common/easy-lease/pkg/models"
	"github.com/gorilla/mux"
)

// GetComplaints godoc
// @Summary Get details of all complaints
// @Description Get details of all complaints
// @Tags complaints
// @Accept  json
// @Produce  json
//@Param token header string true "token header"
// @Success 200 {array} models.Complaint
// @Router /complaints [get]
func GetComplaints(w http.ResponseWriter, r *http.Request) {
	newComplaints := models.GetAllComplaints()
	res, _ := json.Marshal(newComplaints)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetComplaintById godoc
// @Summary Get details of complaint
// @Description Get details of complaint
// @Tags complaints
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param leaseId path string true "complaintId"
// @Success 200 {array} models.Complaint
// @Router /complaints/{complaintId} [get]
func GetComplaintById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	complaintId := vars["complaintId"]
	ID, err := strconv.ParseInt(complaintId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	complaintDetails, _ := models.GetComplaintById(ID)

	res, _ := json.Marshal(complaintDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// CreateComplaint godoc
// @Summary Create complaint
// @Description Create a new complaint
// @Tags complaints
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param body body string true "body"
// @Success 200 {array} models.Complaint
// @Router /complaints [post]
func CreateComplaint(w http.ResponseWriter, r *http.Request) {
	c := &models.Complaint{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	complaint := c.CreateComplaint()
	res, _ := json.Marshal(complaint)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Delete complaint godoc
// @Summary Delete complaint
// @Description Delete the existing complaint
// @Tags complaints
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Success 200 {array} models.Complaint
// @Router /complaints/{complaintId} [delete]
func DeleteComplaint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	complaintId := vars["complaintId"]
	ID, err := strconv.ParseInt(complaintId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	complaint := models.DeleteComplaint(ID)
	res, _ := json.Marshal(complaint)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetComplaintByUser godoc
// @Summary Get all complaints of user
// @Description Get details of complaint made by a specific user
// @Tags complaints
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param leaseId path string true "userId"
// @Success 200 {array} models.Complaint
// @Router /complaints/user/{userId} [get]
func GetComplaintsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	complaints := models.GetComplaintsByUserId(ID)
	res, _ := json.Marshal(complaints)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
