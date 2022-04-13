package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/common/easy-lease/pkg/models"
	"github.com/common/easy-lease/pkg/utils"
	"github.com/gorilla/mux"
)

// GetComplaints godoc
// @Summary Get details of all apartments
// @Description Get details of all apartments
// @Tags apartments
// @Accept  json
// @Produce  json
//@Param token header string true "token header"
// @Success 200 {array} models.Apartment
// @Router /apartments [get]
func GetApartment(w http.ResponseWriter, r *http.Request) {
	newApartments := models.GetAllApartments()
	res, _ := json.Marshal(newApartments)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetApartmentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apartmentId := vars["apartmentId"]
	ID, err := strconv.ParseInt(apartmentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	apartmentDetails, _ := models.GetApartmentById(ID)

	res, _ := json.Marshal(apartmentDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateApartment(w http.ResponseWriter, r *http.Request) {
	createApartment := &models.Apartment{}
	utils.ParseBody(r, createApartment)
	apartment := createApartment.CreateApartment()
	res, _ := json.Marshal(apartment)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteApartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apartmentId := vars["apartmentId"]
	ID, err := strconv.ParseInt(apartmentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	apartment := models.DeleteApartment(ID)
	res, _ := json.Marshal(apartment)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateApartment(w http.ResponseWriter, r *http.Request) {
	var updateApartment = &models.Apartment{}
	utils.ParseBody(r, updateApartment)
	vars := mux.Vars(r)
	apartmentId := vars["apartmentId"]
	ID, err := strconv.ParseInt(apartmentId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	apartmentDetails, db := models.GetApartmentById(ID)
	var apartmentName string = apartmentDetails.Name
	var apartmentAddress string = apartmentDetails.Address
	var apartmentAmenities string = apartmentDetails.Amenities

	if apartmentName != "" {
		apartmentDetails.Name = apartmentName
	}

	if apartmentAddress != "" {
		apartmentDetails.Address = apartmentAddress
	}

	if apartmentAmenities != "" {
		apartmentDetails.Amenities = apartmentAmenities
	}

	db.Save(&apartmentDetails)

	res, _ := json.Marshal(apartmentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
