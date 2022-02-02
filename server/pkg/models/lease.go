package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lease struct {
	gorm.Model
	ListingId      int64  `gorm:""json: "listingId"`
	UserId         int64  `json: "userId"`
	LeaseStartDate string `json: "leaseStartDate"`
	LeaseEndDate   string `json: "leaseEndDate"`
}

func init() {
	config.ConnectSqlite()
	db = config.GetDB()
	db.AutoMigrate(&Lease{})
}

func (l *Lease) CreateLease() *Lease {
	//db.NewRecord(b)

	db.Create(&l)
	return l
}

func GetAllLeases() []Lease {
	var Leases []Lease
	db.Find(&Leases)
	return Leases
}

func GetLeaseById(Id int64) (*Lease, *gorm.DB) {
	var getLease Lease
	db := db.Where("ID=?", Id).Find(&getLease)
	return &getLease, db
}

func DeleteLease(ID int64) *Lease {
	var lease Lease
	db.First(&lease, ID)
	db.Delete(&Lease{}, ID)
	return &lease
}
