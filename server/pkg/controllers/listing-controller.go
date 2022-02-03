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

func GetListing(w http.ResponseWriter, r *http.Request) {
	newListings := models.GetAllListings()
	res, _ := json.Marshal(newListings)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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

func CreateListing(w http.ResponseWriter, r *http.Request) {
	CreateListing := &models.Listing{}
	utils.ParseBody(r, CreateListing)
	listing := CreateListing.CreateListing()
	res, _ := json.Marshal(listing)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

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
