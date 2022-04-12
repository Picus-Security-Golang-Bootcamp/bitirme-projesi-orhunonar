package database

import (
	"crypto/sha256"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

var db *gorm.DB
var err error

func ConnectDatabase() {
	log.Println("Connecting to database...")
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=Bucket password=admin sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

}
func SignIn() {
	user1 := User{
		Name:            "John Doe",
		Username:        "johndoe",
		Email:           "a@b.com",
		Password:        "password",
		ConfirmPassword: "password",
	}
	if err := db.Create(&user1).Error; err != nil {
		panic(err)
	}
	if user1.Password != user1.ConfirmPassword {
		log.Println("Password and Confirm Password do not match")
		return
	}

	sum := sha256.Sum256([]byte(user1.Password))
	sumString := fmt.Sprintf("%x", sum)
	fmt.Println(sumString)

	db.Create(&User{Username: user1.Username, Password: sumString, Email: user1.Email, Name: user1.Name})
	log.Println("User created successfully")

}
