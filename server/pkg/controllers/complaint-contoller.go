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

func GetComplaints(w http.ResponseWriter, r *http.Request) {
	newComplaints := models.GetAllComplaints()
	res, _ := json.Marshal(newComplaints)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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
