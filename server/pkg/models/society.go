package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var sdb *gorm.DB

type Society struct {
	gorm.Model
	SocietyName      string `gorm:""json: "societyName"`
	SocietyAddress   string `json: "societyAddress"`
	SocietyCity      string `json: "societyCity"`
	SocietyAmenities string `json: "societyAmenities"`
	SocietyImg       string `json: "societyImg"`
}

func init() {
	config.ConnectSqlite()
	sdb = config.GetDB()
	sdb.AutoMigrate(&Society{})
}

func (s *Society) CreateSociety() *Society {
	//db.NewRecord(b)
	sdb.Create(&s)
	return s
}

func GetAllSocieties() []Society {
	var Societies []Society
	sdb.Find(&Societies)
	return Societies
}

func GetSocietyById(Id int64) (*Society, *gorm.DB) {
	var getSociety Society
	sdb := sdb.Where("ID=?", Id).Find(&getSociety)
	return &getSociety, sdb
}

func DeleteSociety(ID int64) *Society {
	var society Society
	sdb.First(&society, ID)
	sdb.Delete(&Society{}, ID)
	return &society
}
