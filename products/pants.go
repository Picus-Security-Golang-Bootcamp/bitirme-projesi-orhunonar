package products

import (
	"finalproject/products/categories"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InsertPants() {
	// Open the CSV file for reading
	file, err := os.Open("./csvfiles/pants.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new reader.
	var pants []categories.Pants
	err = gocsv.Unmarshal(file, &pants)
	if err != nil {
		panic(err)
	}

	// Open a postgres database connection using GORM
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create `pants` table if not exists
	err = db.AutoMigrate(&categories.Pants{})
	if err != nil {
		panic(err)
	}

	// Save all the records at once in the database
	result := db.Create(pants)
	if result.Error != nil {
		panic(result.Error)
	}

	// Print the number of records saved
	fmt.Printf("%d pants saved\n", len(pants))
}

func AddPants(Brand string, Name string, Price float64, Stock int, Size int, Rating float64, Gender string) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if CheckPant(Name) {
		fmt.Println("Pant already exists")
	} else {

		err = db.AutoMigrate(&categories.Pants{})
		if err != nil {
			panic(err)
		}

		pants := categories.Shoes{Brand: Brand, Name: Name, Price: Price, Stock: Stock, Size: Size, Rating: Rating}
		db.Create(&pants)

	}
}

func DeletePants(Name string) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var pant categories.Pants
	db.Where("name = ?", Name).First(&pant)
	if pant.Name == Name {
		db.Delete(&pant)
		log.Println("Pant deleted successfully")
	} else {
		log.Println("Pant not found")
	}
}

func CheckPant(Name string) bool {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var pant categories.Pants
	db.Where("name = ?", Name).First(&pant)
	if pant.Name == Name {
		return true
	} else {
		return false
	}
}
func ListPants(c *gin.Context) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var pants []categories.Pants
	db.Find(&pants)
	for _, pant := range pants {
		fmt.Println(pant.Name)
	}
	c.JSON(200, gin.H{
		"pants": pants,
	})
}
func BuyPants(Name string, Stock int) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var pant categories.Pants
	db.Where("name = ?", Name).First(&pant)
	if pant.Name == Name {
		pant.Stock = pant.Stock - Stock
		db.Save(&pant)
		log.Println("Pant bought successfully")
	} else {
		log.Println("Pant not found")
	}
}
