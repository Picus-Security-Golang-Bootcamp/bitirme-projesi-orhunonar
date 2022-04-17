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

var db *gorm.DB

func InsertShoes() {

	// Open the CSV file for reading
	file, err := os.Open("./csvfiles/shoes.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new reader.
	var shoes []categories.Shoes
	err = gocsv.Unmarshal(file, &shoes)
	if err != nil {
		panic(err)
	}

	// Open a postgres database connection using GORM
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create `shoes` table if not exists
	err = db.AutoMigrate(&categories.Shoes{})
	if err != nil {
		panic(err)
	}

	// Save all the records at once in the database
	result := db.Create(shoes)
	if result.Error != nil {
		panic(result.Error)
	}

	// Print the number of records saved
	fmt.Printf("%d shoes saved\n", len(shoes))
}

func AddShoe(Brand string, Name string, Price float64, Stock int, Size int, Rating float64, Gender string) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if CheckShoe(Name) {
		fmt.Println("Shoe already exists")
	} else {

		// Create `shoes` table if not exists
		err = db.AutoMigrate(&categories.Shoes{})
		if err != nil {
			panic(err)
		}

		shoes := categories.Shoes{Brand: Brand, Name: Name, Price: Price, Stock: Stock, Size: Size, Rating: Rating}
		db.Create(&shoes)

	}
}

func DeleteShoes(Name string) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var shoe categories.Shoes
	db.Where("name = ?", Name).First(&shoe)
	if shoe.Name == Name {
		db.Delete(&shoe)
		log.Println("Shoes deleted successfully")
	} else {
		log.Println("Shoes not found")
	}
}
func CheckShoe(Name string) bool {
	var shoe categories.Shoes
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).First(&shoe)
	if shoe.Name == Name {
		return true
	} else {
		return false
	}
}

//List all shoes
// func ListShoes(c *gin.Context) {
// 	var shoes []categories.Shoes
// 	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.Find(&shoes)
// 	for _, shoe := range shoes {
// 		fmt.Println(shoe.Name)
// 	}
// 	c.JSON(200, gin.H{
// 		"shoes": shoes,
// 	})

// }

//List all shoes with pagination
func ListShoes(c *gin.Context, page int, limit int) {
	var shoes []categories.Shoes

	db.Limit(limit).Offset((page - 1) * limit).Find(&shoes)
	for _, shoe := range shoes {
		fmt.Println(shoe.Name)
	}
	c.JSON(200, shoes)

}

func BuyShoes(Name string, Stock int) {
	var shoe categories.Shoes
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).First(&shoe)
	if shoe.Name == Name {
		if shoe.Stock > 0 {
			shoe.Stock = shoe.Stock - Stock
			db.Save(&shoe)
			log.Println("Shoes bought successfully")
		} else {
			log.Println("Out of stock")
		}
	} else {
		log.Println("Shoes not found")
	}
}
func CountShoes() int {
	var shoes []categories.Shoes
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Find(&shoes)
	return len(shoes)
}
