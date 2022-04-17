package productapi

import (
	"finalproject/api"
	"finalproject/pagination"
	"finalproject/products"
	database "finalproject/user_database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePantsCategory(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

		products.InsertPants()

		c.JSON(200, gin.H{
			"message": "Pants category created successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}

func AddPants(c *gin.Context) {

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

		products.AddPants(Brand, Name, NewPrice, NewStock, NewSize, NewRating, Gender)
		c.JSON(200, gin.H{
			"message": "Pants added successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}

func DeletePants(c *gin.Context) {

	if database.VerifyAdmin(api.MyToken) {

		Name := c.Params.ByName("name")
		products.DeletePants(Name)
		c.JSON(200, gin.H{
			"message": "Pants deleted successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}

func ListPants(c *gin.Context) {

	var pagination pagination.Pagination
	page := c.Params.ByName("page")
	limit := c.Params.ByName("limit")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	pagination.SetPagination(pageInt, limitInt, products.CountPants())
	products.ListPants(c, pagination.Page, pagination.Limit)

	c.JSON(200, gin.H{
		"message": "Pants listed successfully",
	})
}

func BuyPants(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Params.ByName("name")
		Stock := c.Params.ByName("stock")

		NewStock, _ := strconv.Atoi(Stock)

		products.BuyPants(Name, NewStock)
		c.JSON(200, gin.H{
			"message": "Pants bought successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}

}
