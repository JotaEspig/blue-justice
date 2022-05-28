package user

import (
	dbPKG "sigma/db"

	"gorm.io/gorm"
)

// Slice of all user params
var UserParams = []string{
	"id",
	"username",
	"name",
	"surname",
	"email",
	"password",
	"type",
}

// Slice of public user params
var PublicUserParams = []string{
	"id",
	"username",
	"name",
	"surname",
	"email",
	"type",
}

// Adds an user to a database.
func AddUser(db *gorm.DB, u *User) error {
	return db.Create(u).Error
}

// Gets an user from a database
func GetUser(db *gorm.DB, username string, columns ...string) (*User, error) {
	u := &User{}

	columnsToUse := dbPKG.GetColumns(UserParams, columns...)

	err := db.Select(columnsToUse).Where("username = ?", username).First(u).Error

	return u, err
}

// Removes an user from a database
func RmUser(db *gorm.DB, username string) error {
	return db.Unscoped().Where("username = ?", username).Delete(&User{}).Error
}

// AutoMigrate the user table
func init() {
	dbPKG.DB.AutoMigrate(&User{})
}
