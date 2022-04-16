package api

import (
	database "finalproject/user_database"

	"github.com/gin-gonic/gin"
)

// Homepage
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Home Page",
	})
}

// Database Connection
func ConnectDatabase(c *gin.Context) {
	database.ConnectDatabase()
	c.JSON(200, gin.H{
		"message": "Connected to database successfully",
	})
}

// SignIn
func SignIn(c *gin.Context) {
	Name := c.Params.ByName("name")
	Username := c.Params.ByName("username")
	Password := c.Params.ByName("password")
	Email := c.Params.ByName("email")
	ConfirmPassword := c.Params.ByName("confirm_password")

	database.SignIn(Name, Username, Email, Password, ConfirmPassword)
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

// Login
func LogIn(c *gin.Context) {
	Username := c.Params.ByName("username")
	Password := c.Params.ByName("password")
	database.LogIn(Username, Password)
	c.JSON(200, gin.H{
		"message": "User logged in successfully",
	})
}

func SuperUser(c *gin.Context) {
	database.CreateSuperUser()
	c.JSON(200, gin.H{
		"message": "Super User logged in successfully",
	})
}

func ListUsers(c *gin.Context) {
	database.ListUsers(c)
	c.JSON(200, gin.H{
		"message": "Users Listed successfully",
	})
}

func DeleteUser(c *gin.Context) {
	Username := c.Params.ByName("username")
	database.DeleteUser(Username)
	c.JSON(200, gin.H{
		"message": "User deleted successfully",
	})
}
func SearchUser(c *gin.Context) {
	Username := c.Params.ByName("username")
	database.SearchUser(c, Username)
	c.JSON(200, gin.H{
		"message": "User searched successfully",
	})
}
