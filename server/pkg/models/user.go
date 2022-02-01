package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email    string `gorm:""json: "email"`
	Name     string `json: "name"`
	Password string `json: "password"`
}

func init() {
	config.ConnectTest()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	//db.NewRecord(b)

	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Email=?", email).Find(&getUser)
	return &getUser, db
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func GetUserByName(name string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("Name=?", name).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
