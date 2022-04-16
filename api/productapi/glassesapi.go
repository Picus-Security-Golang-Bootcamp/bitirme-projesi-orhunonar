package productapi

import (
	"finalproject/api"
	"finalproject/products"
	database "finalproject/user_database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Add Glasses
func CreateGlassesCategory(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

		products.InsertGlasses()

		c.JSON(200, gin.H{
			"message": "Glasses category created successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}

func AddGlasses(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

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
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}

func DeleteGlasses(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

		Name := c.Params.ByName("name")

		products.DeleteGlasses(Name)
		c.JSON(200, gin.H{
			"message": "Glasses deleted successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}

func ListGlasses(c *gin.Context) {

	products.ListGlasses(c)
	c.JSON(200, gin.H{
		"message": "Glasses listed successfully",
	})

}

func BuyGlasses(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Params.ByName("name")
		Stock := c.Params.ByName("stock")

		NewStock, _ := strconv.Atoi(Stock)

		products.BuyGlasses(Name, NewStock)
		c.JSON(200, gin.H{
			"message": "Glasses bought successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}
