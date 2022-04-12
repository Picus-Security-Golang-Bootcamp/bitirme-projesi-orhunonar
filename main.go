package main

import (
	"finalproject/database"

	"gorm.io/gorm"
)

var err error
var db *gorm.DB

func main() {

	database.ConnectDatabase()
	database.CreateSuperUser()

}
