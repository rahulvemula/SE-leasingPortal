package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var ldb *gorm.DB

type Listing struct {
	gorm.Model
	ListingType string `gorm:""json:"listingType"`
	HouseType   string `json:"houseType"`
	Rent        int64  `json: "rent"`
	UserId      int64  `json: "userId"`
	Isleased    bool   `gorm:"default:false" json: "isleased"`
	ListingImg  string `json: "listingImg"`
	SocietyId   string `json: "societyId"`
}

func init() {
	config.ConnectSqlite()
	ldb = config.GetDB()
	ldb.AutoMigrate(&Listing{})
}

func (l *Listing) CreateListing() *Listing {
	db.Create(&l)
	return l
}

func GetAllListings() []Listing {
	var Listings []Listing
	db.Find(&Listings)
	return Listings
}

func GetListingById(Id int64) (*Listing, *gorm.DB) {
	var getListing Listing
	db := db.Where("ID=?", Id).Find(&getListing)
	return &getListing, db
}

func DeleteListing(ID int64) *Listing {
	var listing Listing
	db.First(&listing, ID)
	db.Delete(&Listing{}, ID)
	return &listing
}
