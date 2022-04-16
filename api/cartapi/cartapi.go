package cartapi

import (
	"finalproject/products/cart"
	"finalproject/products/oldorders"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddShoeToCart(c *gin.Context) {
	Name := c.Param("name")
	Stock := c.Param("stock")

	NewStock, _ := strconv.Atoi(Stock)

	cart.AddShoeToCart(Name, NewStock)

	c.JSON(200, gin.H{
		"message": "Shoe added to cart successfully",
	})
}
func AddGlassesToCart(c *gin.Context) {
	Name := c.Param("name")
	Stock := c.Param("stock")

	NewStock, _ := strconv.Atoi(Stock)

	cart.AddGlassesToCart(Name, NewStock)

	c.JSON(200, gin.H{
		"message": "Glasses added to cart successfully",
	})
}
func AddPantsToCart(c *gin.Context) {
	Name := c.Param("name")
	Stock := c.Param("stock")

	NewStock, _ := strconv.Atoi(Stock)

	cart.AddPantsToCart(Name, NewStock)

	c.JSON(200, gin.H{
		"message": "Pants added to cart successfully",
	})
}
func RemoveProductFromCart(c *gin.Context) {
	Name := c.Param("name")

	cart.RemoveProductFromCart(Name)

	c.JSON(200, gin.H{
		"message": "Product removed from cart successfully",
	})
}
func UpdateProductInCart(c *gin.Context) {
	Name := c.Param("name")
	Stock := c.Param("stock")

	NewStock, _ := strconv.Atoi(Stock)

	cart.UpdateProductStock(Name, NewStock)

	c.JSON(200, gin.H{
		"message": "Product updated in cart successfully",
	})
}
func ListCart(c *gin.Context) {

	cart.ListProducts(c)

	c.JSON(200, gin.H{
		"message": "Cart listed successfully",
	})
}
func GetTotalPrice(c *gin.Context) {

	cart.GetTotalPrice(c)

	c.JSON(200, gin.H{
		"message": "Total price listed successfully",
	})
}
func CompleteOrder(c *gin.Context) {

	cart.CompleteOrder()

	c.JSON(200, gin.H{
		"message": "Order completed successfully",
	})
}
func ClearCart(c *gin.Context) {

	cart.ClearCart()

	c.JSON(200, gin.H{
		"message": "Cart cleared successfully",
	})
}
func CancelOrder(c *gin.Context) {

	cart.CancelOrder(c)

	c.JSON(200, gin.H{
		"message": "Order cancelled successfully",
	})
}

func GetOldOrders(c *gin.Context) {

	oldorders.GetOldOrders(c)

	c.JSON(200, gin.H{
		"message": "Old orders listed successfully",
	})
}
func SearchProductInCart(c *gin.Context) {

	Name := c.Param("name")

	cart.SearchProductInCart(Name, c)

	c.JSON(200, gin.H{
		"message": "Product searched in cart successfully",
	})
}
