package productapi

import (
	"finalproject/products"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePantsCategory(c *gin.Context) {

	products.InsertPants()

	c.JSON(200, gin.H{
		"message": "Pants category created successfully",
	})
}

func AddPants(c *gin.Context) {
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

	products.AddPants(Brand, Name, NewPrice, NewStock, NewSize, NewRating, Gender)
	c.JSON(200, gin.H{
		"message": "Pants added successfully",
	})
}

func DeletePants(c *gin.Context) {
	Name := c.Params.ByName("name")
	products.DeletePants(Name)
	c.JSON(200, gin.H{
		"message": "Pants deleted successfully",
	})

}

func ListPants(c *gin.Context) {

	products.ListPants(c)
	c.JSON(200, gin.H{
		"message": "Pants listed successfully",
	})

}

func BuyPants(c *gin.Context) {
	Name := c.Params.ByName("name")
	Stock := c.Params.ByName("stock")

	NewStock, _ := strconv.Atoi(Stock)

	products.BuyPants(Name, NewStock)
	c.JSON(200, gin.H{
		"message": "Pants bought successfully",
	})

}
