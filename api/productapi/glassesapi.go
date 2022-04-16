package productapi

import (
	"finalproject/products"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add Glasses
func CreateGlassesCategory(c *gin.Context) {

	products.InsertGlasses()

	c.JSON(200, gin.H{
		"message": "Glasses category created successfully",
	})
}

func AddGlasses(c *gin.Context) {

	Brand := c.Params.ByName("brand")
	Name := c.Params.ByName("name")
	Price := c.Params.ByName("price")
	Stock := c.Params.ByName("stock")
	Size := c.Params.ByName("size")
	Rating := c.Params.ByName("rating")
	Gender := c.Params.ByName("gender")

	NewPrice, _ := strconv.ParseFloat(Price, 64)
	NewStock, _ := strconv.Atoi(Stock)
	NewSize, _ := strconv.Atoi(Size)
	NewRating, _ := strconv.ParseFloat(Rating, 64)

	products.AddGlasses(Brand, Name, NewPrice, NewStock, NewSize, NewRating, Gender)

	c.JSON(200, gin.H{
		"message": "Glasses added successfully",
	})
}

func DeleteGlasses(c *gin.Context) {
	Name := c.Params.ByName("name")

	products.DeleteGlasses(Name)
	c.JSON(200, gin.H{
		"message": "Glasses deleted successfully",
	})

}

func ListGlasses(c *gin.Context) {

	products.ListGlasses(c)
	c.JSON(200, gin.H{
		"message": "Glasses listed successfully",
	})

}

func BuyGlasses(c *gin.Context) {
	Name := c.Params.ByName("name")
	Stock := c.Params.ByName("stock")

	NewStock, _ := strconv.Atoi(Stock)

	products.BuyGlasses(Name, NewStock)
	c.JSON(200, gin.H{
		"message": "Glasses bought successfully",
	})

}
