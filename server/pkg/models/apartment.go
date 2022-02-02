package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var apdb *gorm.DB

type Apartment struct {
	gorm.Model
	Name      string `gorm:""json: "name"`
	Address   string `json: "address"`
	Amenities string `json: "amenities"`
}

func init() {
	config.ConnectSqlite()
	apdb = config.GetDB()
	apdb.AutoMigrate(&Apartment{})
}

func (a *Apartment) CreateApartment() *Apartment {
	apdb.Create(&a)
	return a
}

func GetAllApartments() []Apartment {
	var Apartments []Apartment
	apdb.Find(&Apartments)
	return Apartments
}

func GetApartmentById(Id int64) (*Apartment, *gorm.DB) {
	var getApartment Apartment
	apdb := apdb.Where("ID=?", Id).Find(&getApartment)
	return &getApartment, apdb
}

func DeleteApartment(ID int64) *Apartment {
	var apartment Apartment
	apdb.First(&apartment, ID)
	apdb.Delete(&Apartment{}, ID)
	return &apartment
}
