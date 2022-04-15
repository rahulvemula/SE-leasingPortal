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

// GetLeases godoc
// @Summary Get details of all leases
// @Description Get details of all leases
// @Tags leases
// @Accept  json
// @Produce  json
//@Param token header string true "token header"
// @Success 200 {array} models.Lease
// @Router /leases [get]
func GetLease(w http.ResponseWriter, r *http.Request) {
	newLeases := models.GetAllLeases()
	res, _ := json.Marshal(newLeases)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetLeaseById godoc
// @Summary Get details lease
// @Description Get details of leases
// @Tags leases
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param leaseId path string true "leaseId"
// @Success 200 {array} models.Lease
// @Router /leases/{leaseId} [get]
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

// CreateLease godoc
// @Summary Create lease
// @Description Create a new lease
// @Tags leases
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param body body string true "body"
// @Success 200 {array} models.Lease
// @Router /leases [post]
func CreateLease(w http.ResponseWriter, r *http.Request) {
	CreateLease := &models.Lease{}
	utils.ParseBody(r, CreateLease)
	lease := CreateLease.CreateLease()
	res, _ := json.Marshal(lease)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Delete lease godoc
// @Summary Delete lease
// @Description Delete the existing lease
// @Tags leases
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Success 200 {array} models.Lease
// @Router /leases/{leaseId} [delete]
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

// Update lease godoc
// @Summary Update lease
// @Description Update the existing lease
// @Tags leases
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param leaseId path string true "leaseId"
// @Param body body string true "body"
// @Success 200 {array} models.Lease
// @Router /leases/{leaseId} [put]
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
