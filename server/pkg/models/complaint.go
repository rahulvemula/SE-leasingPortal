package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var cdb *gorm.DB

type Complaint struct {
	gorm.Model
	UserId      int64  `gorm:""json: "userId"`
	Description string `json: "description"`
}

func init() {
	config.ConnectSqlite()
	cdb = config.GetDB()
	cdb.AutoMigrate(&Complaint{})
}

func (a *Complaint) CreateComplaint() *Complaint {
	cdb.Create(&a)
	return a
}

func GetAllComplaints() []Complaint {
	var Complaints []Complaint
	cdb.Find(&Complaints)
	return Complaints
}

func GetComplaintById(Id int64) (*Complaint, *gorm.DB) {
	var getComplaint Complaint
	cdb := cdb.Where("ID=?", Id).Find(&getComplaint)
	return &getComplaint, cdb
}

func GetComplaintsByUserId(Id int64) []Complaint {
	var Complaints []Complaint
	cdb.Where("user_id = ?", Id).Find(&Complaints)
	return Complaints
}

func DeleteComplaint(ID int64) *Complaint {
	var complaint Complaint
	cdb.First(&complaint, ID)
	cdb.Delete(&Complaint{}, ID)
	return &complaint
}
