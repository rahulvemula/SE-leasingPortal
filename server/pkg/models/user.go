package models

import (
	"github.com/common/easy-lease/pkg/config"
	"gorm.io/gorm"
)

var userdb *gorm.DB

type User struct {
	gorm.Model
	Email    string `gorm:""json: "email"`
	Name     string `json: "name"`
	Password string `json: "password"`
}

func init() {
	config.ConnectSqlite()
	userdb = config.GetDB()
	userdb.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	//userdb.NewRecord(b)

	userdb.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	userdb.Find(&Users)
	return Users
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var getUser User
	userdb := userdb.Where("Email=?", email).Find(&getUser)
	return &getUser, userdb
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	userdb := userdb.Where("ID=?", Id).Find(&getUser)
	return &getUser, userdb
}

func GetUserByName(name string) (*User, *gorm.DB) {
	var getUser User
	userdb := userdb.Where("Name=?", name).Find(&getUser)
	return &getUser, userdb
}

func DeleteUser(ID int64) *User {
	var user User
	userdb.First(&user, ID)
	userdb.Delete(&User{}, ID)
	return &user
}
