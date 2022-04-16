package handle

import "github.com/gin-gonic/gin"

//Not Found error
func NotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Not Found",
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"message": message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(500, gin.H{
		"message": message,
	})
}

func Created(c *gin.Context, message string) {
	c.JSON(201, gin.H{
		"message": message,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(204, gin.H{
		"message": "No Content",
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(401, gin.H{
		"message": "Unauthorized",
	})
}

func Forbidden(c *gin.Context) {
	c.JSON(403, gin.H{
		"message": "Forbidden",
	})
}

func Conflict(c *gin.Context, message string) {
	c.JSON(409, gin.H{
		"message": message,
	})
}

func NotImplemented(c *gin.Context) {
	c.JSON(501, gin.H{
		"message": "Not Implemented",
	})
}

func NotAcceptable(c *gin.Context) {
	c.JSON(406, gin.H{
		"message": "Not Acceptable",
	})
}
