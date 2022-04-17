package main

import (
	"finalproject/api"
	"finalproject/api/cartapi"
	"finalproject/api/productapi"
	"finalproject/products/categories"
	"finalproject/products/search"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB
var c *gin.Context

func main() {

	r := gin.Default()

	//Initials
	r.POST("/", api.ConnectDatabase)
	r.POST("/homepage", api.HomePage)

	// FOR USER
	r.POST("/SignIn/:name,username,password,email,confirm_password", api.SignIn)
	r.POST("/LogIn/:username,password", api.LogIn)
	r.POST("/create_super_user", api.SuperUser)
	r.GET("/list_users/:page,limit", api.ListUsers)
	r.DELETE("/delete_user/:username", api.DeleteUser)
	r.GET("/get_user/:username", api.SearchUser)

	// FOR SHOES
	r.POST("/create_shoes_category", productapi.CreateShoesCategory)
	r.POST("/add_shoes/:brand,name,price,stock,rating,size,", productapi.AddShoes)
	r.GET("/list_shoes/:page,limit", productapi.ListShoes)
	r.DELETE("/delete_shoes/:name", productapi.DeleteShoes)
	r.GET("/get_shoes/:name", search.SearchShoesByName)
	r.PATCH("/buy_shoes/:name,stock", productapi.BuyShoes)

	// FOR PANTS
	r.POST("/create_pants_category", productapi.CreatePantsCategory)
	r.POST("/add_pants/:brand,name,price,stock,rating,size,", productapi.AddPants)
	r.GET("/list_pants/:page,limit", productapi.ListPants)
	r.DELETE("/delete_pants/:name", productapi.DeletePants)
	r.GET("/get_pants/:name", search.SearchPantsByName)
	r.PATCH("/buy_pants/:name,stock", productapi.BuyPants)

	// FOR GLASSES
	r.POST("/create_glasses_category", productapi.CreateGlassesCategory)
	r.POST("/add_glasses/:brand,name,price,stock,rating,size,", productapi.AddGlasses)
	r.GET("/list_glasses/:page,limit", productapi.ListGlasses)
	r.DELETE("/delete_glasses/:name", productapi.DeleteGlasses)
	r.GET("/get_glasses/:name", search.SearchGlassesByName)
	r.PATCH("/buy_glasses/:name,stock", productapi.BuyGlasses)

	//FOR CART
	r.POST("/add_glasses_to_cart/:name,stock", cartapi.AddGlassesToCart)
	r.POST("/add_pants_to_cart/:name,stock", cartapi.AddPantsToCart)
	r.POST("/add_shoes_to_cart/:name,stock", cartapi.AddShoeToCart)
	r.GET("/list_products_in_cart", cartapi.ListCart)
	r.DELETE("/remove_product_from_cart/:name", cartapi.RemoveProductFromCart)
	r.GET("/get_product_in_cart/:name", cartapi.SearchProductInCart)
	r.PATCH("/buy_cart", cartapi.UpdateProductInCart)
	r.GET("/get_total_price", cartapi.GetTotalPrice)
	r.POST("/checkout", cartapi.CompleteOrder)
	r.DELETE("/delete_cart", cartapi.ClearCart)
	r.POST("/cancled_order", cartapi.CancelOrder)
	r.GET("/old_order", cartapi.GetOldOrders)

	//List Categories
	r.GET("/list_categories", categories.ListCategories)

	// listen and serve on port 8080
	r.Run(":8080")

}
