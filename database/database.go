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

// connect to database
func ConnectDatabase() {
	log.Println("Connecting to database...")
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=Bucket password=admin sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	log.Println("Connected to database successfully")
}

// sign in
func SignIn(Name string, Username string, Email string, Password string, ConfirmPassword string) {

	if Password != ConfirmPassword {
		fmt.Println("Password and Confirm Password do not match")
	} else if CheckUsername(Username) {
		fmt.Println("Username already used")
	} else if CheckEmail(Email) {
		fmt.Println("Email already used")
	} else {
		sum := sha256.Sum256([]byte(Password))
		sumString := fmt.Sprintf("%x", sum)
		user := User{Name: Name, Username: Username, Email: Email, Password: sumString}
		db.Create(&user)
		fmt.Println("User created successfully")

	}

}

// check username is already used
func CheckUsername(Username string) bool {
	var user User
	db.Where("username = ?", Username).First(&user)
	if user.Username == Username {
		return true
	} else {

		return false
	}
}

//check e-mail is already used
func CheckEmail(Email string) bool {
	var user User
	db.Where("email = ?", Email).First(&user)
	if user.Email == Email {

		return true
	} else {

		return false
	}
}

// log in
func LogIn(Username string, Password string) bool {
	var user User
	db.Where("username = ?", Username).First(&user)
	if user.Username == Username {
		sum := sha256.Sum256([]byte(Password))
		sumString := fmt.Sprintf("%x", sum)
		if sumString == user.Password {
			log.Println("Logged in successfully")
			return true
		} else {
			log.Println("Wrong password")
			return false
		}
	} else {
		log.Println("Wrong username")
		return false
	}
}

// create super user
func CreateSuperUser() {
	var user User
	db.Where("username = ?", "admin").First(&user)
	if user.Username == "admin" {
		log.Println("Super user already created")
	} else {
		sum := sha256.Sum256([]byte("admin"))
		sumString := fmt.Sprintf("%x", sum)
		user := User{Name: "admin", Username: "admin", Email: "admin@gmail.com", Password: sumString}
		db.Create(&user)
		log.Println("Super user created successfully")
	}
}
