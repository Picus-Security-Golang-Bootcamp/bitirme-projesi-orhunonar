package cartapi

import (
	"finalproject/api"
	"finalproject/products/cart"
	"finalproject/products/oldorders"
	database "finalproject/user_database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddShoeToCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Param("name")
		Stock := c.Param("stock")

		NewStock, _ := strconv.Atoi(Stock)

		cart.AddShoeToCart(Name, NewStock)

		c.JSON(200, gin.H{
			"message": "Shoe added to cart successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func AddGlassesToCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Param("name")
		Stock := c.Param("stock")

		NewStock, _ := strconv.Atoi(Stock)

		cart.AddGlassesToCart(Name, NewStock)

		c.JSON(200, gin.H{
			"message": "Glasses added to cart successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func AddPantsToCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Param("name")
		Stock := c.Param("stock")

		NewStock, _ := strconv.Atoi(Stock)

		cart.AddPantsToCart(Name, NewStock)

		c.JSON(200, gin.H{
			"message": "Pants added to cart successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func RemoveProductFromCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Param("name")

		cart.RemoveProductFromCart(Name)

		c.JSON(200, gin.H{
			"message": "Product removed from cart successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func UpdateProductInCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Param("name")
		Stock := c.Param("stock")

		NewStock, _ := strconv.Atoi(Stock)

		cart.UpdateProductStock(Name, NewStock)

		c.JSON(200, gin.H{
			"message": "Product updated in cart successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func ListCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		cart.ListProducts(c)

		c.JSON(200, gin.H{
			"message": "Cart listed successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func GetTotalPrice(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		cart.GetTotalPrice(c)

		c.JSON(200, gin.H{
			"message": "Total price listed successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func CompleteOrder(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		cart.CompleteOrder()

		c.JSON(200, gin.H{
			"message": "Order completed successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func ClearCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		cart.ClearCart()

		c.JSON(200, gin.H{
			"message": "Cart cleared successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func CancelOrder(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		cart.CancelOrder(c)

		c.JSON(200, gin.H{
			"message": "Order cancelled successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}

func GetOldOrders(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		oldorders.GetOldOrders(c)

		c.JSON(200, gin.H{
			"message": "Old orders listed successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
func SearchProductInCart(c *gin.Context) {

	if database.VerifyToken(api.MyToken) {

		Name := c.Param("name")

		cart.SearchProductInCart(Name, c)

		c.JSON(200, gin.H{
			"message": "Product searched in cart successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
