package products

import (
	"finalproject/products/categories"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InsertGlasses() {
	// Open the CSV file for reading
	file, err := os.Open("./csvfiles/glasses.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new reader.
	var glasses []categories.Glasses
	err = gocsv.Unmarshal(file, &glasses)
	if err != nil {
		panic(err)
	}

	// Open a postgres database connection using GORM
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create `glasses` table if not exists
	err = db.AutoMigrate(&categories.Glasses{})
	if err != nil {
		panic(err)
	}

	// Save all the records at once in the database
	result := db.Create(glasses)
	if result.Error != nil {
		panic(result.Error)
	}

	// Print the number of records saved
	fmt.Printf("%d glasses saved\n", len(glasses))
}
func AddGlasses(Brand string, Name string, Price float64, Stock int, Size int, Rating float64, Gender string) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if CheckGlasses(Name) {
		fmt.Println("Glasses already exists")
	} else {

		err = db.AutoMigrate(&categories.Glasses{})
		if err != nil {
			panic(err)
		}

		glasses := categories.Glasses{Brand: Brand, Name: Name, Price: Price, Stock: Stock, Size: Size, Rating: Rating}
		db.Create(&glasses)

	}
}

func DeleteGlasses(Name string) {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var glasses categories.Glasses
	db.Where("name = ?", Name).First(&glasses)
	if glasses.Name == Name {
		db.Delete(&glasses)
		log.Println("Glasses deleted successfully")
	} else {
		log.Println("Glasses not found")
	}
}

func CheckGlasses(Name string) bool {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var glasses categories.Glasses
	db.Where("name = ?", Name).First(&glasses)
	if glasses.Name == Name {
		return true
	} else {
		return false
	}
}

func ListGlasses() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var glasses []categories.Glasses
	db.Find(&glasses)
	for _, glasses := range glasses {
		fmt.Println(glasses)
	}
}
