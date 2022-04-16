package productapi

import (
	"finalproject/api"
	"finalproject/products"
	database "finalproject/user_database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateShoesCategory(c *gin.Context) {
	if database.VerifyAdmin(api.MyToken) {

		products.InsertShoes()

		c.JSON(200, gin.H{
			"message": "Shoes category created successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}

func AddShoes(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

		Brand := c.Params.ByName("brand")
		Name := c.Params.ByName("name")
		Price := c.Params.ByName("price")
		Stock := c.Params.ByName("stock")
		Rating := c.Params.ByName("rating")
		Size := c.Params.ByName("size")
		Gender := c.Params.ByName("gender")

		NewPrice, _ := strconv.ParseFloat(Price, 64)
		NewStock, _ := strconv.Atoi(Stock)
		NewSize, _ := strconv.Atoi(Size)
		NewRating, _ := strconv.ParseFloat(Rating, 64)

		products.AddShoe(Brand, Name, NewPrice, NewStock, NewSize, NewRating, Gender)
		c.JSON(200, gin.H{
			"message": "Shoes added successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}

func DeleteShoes(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

		Name := c.Params.ByName("name")
		products.DeleteShoes(Name)
		c.JSON(200, gin.H{
			"message": "Shoes deleted successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}

func ListShoes(c *gin.Context) {

	products.ListShoes(c)
	c.JSON(200, gin.H{
		"message": "Shoes listed successfully",
	})

}

func BuyShoes(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {
		Name := c.Params.ByName("name")
		Stock := c.Params.ByName("stock")

		NewStock, _ := strconv.Atoi(Stock)
		products.BuyShoes(Name, NewStock)
		c.JSON(200, gin.H{
			"message": "Shoes bought successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
