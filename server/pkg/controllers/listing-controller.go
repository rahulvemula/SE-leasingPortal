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

// GetListing godoc
// @Summary Get details of all listings
// @Description Get details of all listings
// @Tags listings
// @Accept  json
// @Produce  json
//@Param token header string true "token header"
// @Success 200 {array} models.Listing
// @Router /listings [get]
func GetListing(w http.ResponseWriter, r *http.Request) {
	newListings := models.GetAllListings()
	res, _ := json.Marshal(newListings)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetListingById godoc
// @Summary Get details of listing
// @Description Get details of the listing
// @Tags listings
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param leaseId path string true "listingId"
// @Success 200 {array} models.Listing
// @Router /listings/{listingId} [get]
func GetListingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listingId := vars["listingId"]
	ID, err := strconv.ParseInt(listingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	listingDetails, _ := models.GetListingById(ID)

	res, _ := json.Marshal(listingDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// CreateListing godoc
// @Summary Create listings
// @Description Create a new listing
// @Tags listings
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param body body string true "body"
// @Success 200 {array} models.Listing
// @Router /listings [post]
func CreateListing(w http.ResponseWriter, r *http.Request) {
	CreateListing := &models.Listing{}
	utils.ParseBody(r, CreateListing)
	listing := CreateListing.CreateListing()
	res, _ := json.Marshal(listing)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// Delete listing godoc
// @Summary Delete listings
// @Description Delete the existing lease
// @Tags listings
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Success 200 {array} models.Listing
// @Router /listings/{listingId} [delete]
func DeleteListing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listingId := vars["listingId"]
	ID, err := strconv.ParseInt(listingId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	listing := models.DeleteListing(ID)
	res, _ := json.Marshal(listing)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Update listing godoc
// @Summary Update listing
// @Description Update the existing listing
// @Tags listings
// @Accept  json
// @Produce  json
// @Param token header string true "token header"
// @Param body body string true "body"
// @Success 200 {array} models.Listing
// @Router /listings/{listingId} [put]
func UpdateListing(w http.ResponseWriter, r *http.Request) {
	var updateListing = &models.Listing{}
	utils.ParseBody(r, updateListing)
	vars := mux.Vars(r)
	listingId := vars["listingId"]
	ID, err := strconv.ParseInt(listingId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	listingDetails, db := models.GetListingById(ID)
	var listingType string = updateListing.ListingType
	var societyId string = updateListing.SocietyId
	var listingImg string = updateListing.ListingImg
	var houseType string = updateListing.HouseType
	var rent int64 = updateListing.Rent
	var userid int64 = updateListing.UserId
	var isleased bool = updateListing.Isleased

	if listingType != "" {
		listingDetails.ListingType = listingType
	}

	if houseType != "" {
		listingDetails.HouseType = houseType
	}
	if societyId != "" {
		listingDetails.SocietyId = societyId
	}

	if listingImg != "" {
		listingDetails.ListingImg = listingImg
	}
	if rent != 0 {
		listingDetails.Rent = rent
	}

	if userid != 0 {
		listingDetails.UserId = userid
	}

	listingDetails.Isleased = isleased

	db.Save(&listingDetails)

	res, _ := json.Marshal(listingDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
