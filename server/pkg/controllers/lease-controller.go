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

func GetLease(w http.ResponseWriter, r *http.Request) {
	newLeases := models.GetAllLeases()
	res, _ := json.Marshal(newLeases)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLeaseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leaseId := vars["leaseId"]
	ID, err := strconv.ParseInt(leaseId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	leaseDetails, _ := models.GetLeaseById(ID)

	res, _ := json.Marshal(leaseDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateLease(w http.ResponseWriter, r *http.Request) {
	CreateLease := &models.Lease{}
	utils.ParseBody(r, CreateLease)
	lease := CreateLease.CreateLease()
	res, _ := json.Marshal(lease)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteLease(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leaseId := vars["leaseId"]
	ID, err := strconv.ParseInt(leaseId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lease := models.DeleteLease(ID)
	res, _ := json.Marshal(lease)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLease(w http.ResponseWriter, r *http.Request) {
	var updateLease = &models.Lease{}
	utils.ParseBody(r, updateLease)
	vars := mux.Vars(r)
	leaseId := vars["leaseId"]
	ID, err := strconv.ParseInt(leaseId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	leaseDetails, db := models.GetLeaseById(ID)
	var listingId int64 = updateLease.ListingId
	var userId int64 = updateLease.UserId
	var leaseStartDate string = updateLease.LeaseStartDate
	var leaseEndDate string = updateLease.LeaseEndDate

	if listingId != 0 {
		leaseDetails.ListingId = listingId
	}

	if userId != 0 {
		leaseDetails.UserId = userId
	}

	if leaseStartDate != "" {
		leaseDetails.LeaseStartDate = leaseStartDate
	}

	if leaseEndDate != "" {
		leaseDetails.LeaseEndDate = leaseEndDate
	}

	db.Save(&leaseDetails)

	res, _ := json.Marshal(leaseDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
