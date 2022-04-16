package productapi

import (
	"finalproject/products"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateShoesCategory(c *gin.Context) {

	products.InsertShoes()

	c.JSON(200, gin.H{
		"message": "Shoes category created successfully",
	})
}

func AddShoes(c *gin.Context) {
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
}

func DeleteShoes(c *gin.Context) {
	Name := c.Params.ByName("name")
	products.DeleteShoes(Name)
	c.JSON(200, gin.H{
		"message": "Shoes deleted successfully",
	})

}

func ListShoes(c *gin.Context) {

	products.ListShoes(c)
	c.JSON(200, gin.H{
		"message": "Shoes listed successfully",
	})

}

func BuyShoes(c *gin.Context) {
	Name := c.Params.ByName("name")
	Stock := c.Params.ByName("stock")

	NewStock, _ := strconv.Atoi(Stock)
	products.BuyShoes(Name, NewStock)
	c.JSON(200, gin.H{
		"message": "Shoes bought successfully",
	})

}
