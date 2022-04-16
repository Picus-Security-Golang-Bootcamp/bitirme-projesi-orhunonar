package search

import (
	"finalproject/products/categories"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SearchShoesByName(c *gin.Context) {
	Name := c.Params.ByName("name")

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&categories.Shoes{})
	if err != nil {
		panic(err)
	}
	var shoes []categories.Shoes
	db.Where("name = ?", Name).Find(&shoes)
	if len(shoes) > 0 {

		//List details
		fmt.Println("Brand:", shoes[0].Brand)
		fmt.Println("Name:", shoes[0].Name)
		fmt.Println("Price:", shoes[0].Price)
		fmt.Println("Stock:", shoes[0].Stock)
		fmt.Println("Size:", shoes[0].Size)
		fmt.Println("Rating:", shoes[0].Rating)
		fmt.Println("Gender:", shoes[0].Gender)
	} else {
		fmt.Println("Shoe not found")
	}
	c.JSON(200, gin.H{
		"Brand":  shoes[0].Brand,
		"Name":   shoes[0].Name,
		"Price":  shoes[0].Price,
		"Stock":  shoes[0].Stock,
		"Size":   shoes[0].Size,
		"Rating": shoes[0].Rating,
		"Gender": shoes[0].Gender,
	})
}
func SearchPantsByName(c *gin.Context) {
	Name := c.Params.ByName("name")
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&categories.Shoes{})
	if err != nil {
		panic(err)
	}
	var pants []categories.Pants
	db.Where("name = ?", Name).Find(&pants)
	if len(pants) > 0 {

		//List details
		fmt.Println("Brand:", pants[0].Brand)
		fmt.Println("Name:", pants[0].Name)
		fmt.Println("Price:", pants[0].Price)
		fmt.Println("Stock:", pants[0].Stock)
		fmt.Println("Size:", pants[0].Size)
		fmt.Println("Rating:", pants[0].Rating)
		fmt.Println("Gender:", pants[0].Gender)
	} else {
		fmt.Println("Pant not found")
	}
	c.JSON(200, gin.H{
		"Brand":  pants[0].Brand,
		"Name":   pants[0].Name,
		"Price":  pants[0].Price,
		"Stock":  pants[0].Stock,
		"Size":   pants[0].Size,
		"Rating": pants[0].Rating,
		"Gender": pants[0].Gender,
	})
}

func SearchGlassesByName(c *gin.Context) {

	Name := c.Params.ByName("name")
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)

	}

	err = db.AutoMigrate(&categories.Shoes{})
	if err != nil {
		panic(err)
	}
	var glasses []categories.Glasses
	db.Where("name = ?", Name).Find(&glasses)
	if len(glasses) > 0 {

		//List details
		fmt.Println("Brand:", glasses[0].Brand)
		fmt.Println("Name:", glasses[0].Name)
		fmt.Println("Price:", glasses[0].Price)
		fmt.Println("Stock:", glasses[0].Stock)
		fmt.Println("Size:", glasses[0].Size)
		fmt.Println("Rating:", glasses[0].Rating)
		fmt.Println("Gender:", glasses[0].Gender)
	} else {
		fmt.Println("Glasses not found")
	}
	c.JSON(200, gin.H{
		"Brand":  glasses[0].Brand,
		"Name":   glasses[0].Name,
		"Price":  glasses[0].Price,
		"Stock":  glasses[0].Stock,
		"Size":   glasses[0].Size,
		"Rating": glasses[0].Rating,
		"Gender": glasses[0].Gender,
	})
}
