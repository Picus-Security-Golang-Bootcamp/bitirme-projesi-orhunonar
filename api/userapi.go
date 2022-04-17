package api

import (
	"finalproject/pagination"
	database "finalproject/user_database"
	"strconv"

	"github.com/gin-gonic/gin"
)

var MyToken string
var err error

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
	if database.LogIn(Username, Password) {
		MyToken, err = database.GenerateJWT(Username)
		c.JSON(200, gin.H{
			"message": "User logged in successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Wrong password or username",
		})
	}
}

func SuperUser(c *gin.Context) {
	database.CreateSuperUser()
	c.JSON(200, gin.H{
		"message": "Super User logged in successfully",
	})
}

func ListUsers(c *gin.Context) {
	if database.VerifyAdmin(MyToken) {

		// Make pagination when listing users
		var pagination pagination.Pagination
		page := c.Params.ByName("page")
		limit := c.Params.ByName("limit")

		pageInt, _ := strconv.Atoi(page)
		limitInt, _ := strconv.Atoi(limit)

		pagination.SetPagination(pageInt, limitInt, database.CountUsers())
		//List users with pagination
		database.ListUsers(c, pagination.Limit, pagination.Page)

		c.JSON(200, gin.H{
			"message": "Users Listed successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}

func DeleteUser(c *gin.Context) {

	if database.VerifyAdmin(MyToken) {

		Username := c.Params.ByName("username")
		database.DeleteUser(Username)
		c.JSON(200, gin.H{
			"message": "User deleted successfully",
		})

	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to delete users",
		})
	}

}
func SearchUser(c *gin.Context) {
	if database.VerifyAdmin(MyToken) {
		Username := c.Params.ByName("username")
		database.SearchUser(c, Username)
		c.JSON(200, gin.H{
			"message": "User searched successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "You are not authorized to perform this action",
		})
	}
}
