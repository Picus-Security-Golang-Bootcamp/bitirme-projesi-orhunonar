package categories

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Shoes struct {
	gorm.Model
	Brand  string  `csv:"brand"`
	Name   string  `csv:"name"`
	Price  float64 `csv:"price"`
	Stock  int     `csv:"stock"`
	Size   int     `csv:"size"`
	Rating float64 `csv:"rating"`
	Gender string  `csv:"gender"`
}
type Glasses struct {
	gorm.Model
	Brand  string  `csv:"brand"`
	Name   string  `csv:"name"`
	Price  float64 `csv:"price"`
	Stock  int     `csv:"stock"`
	Size   int     `csv:"size"`
	Rating float64 `csv:"rating"`
	Gender string  `csv:"gender"`
}

type Pants struct {
	gorm.Model
	Brand  string  `csv:"brand"`
	Name   string  `csv:"name"`
	Price  float64 `csv:"price"`
	Stock  int     `csv:"stock"`
	Size   int     `csv:"size"`
	Rating float64 `csv:"rating"`
	Gender string  `csv:"gender"`
}

func ListCategories(c *gin.Context) {
	categories := [...]string{"shoes", "glasses", "pants"}
	for _, category := range categories {
		fmt.Println(category)
	}
	c.JSON(200, gin.H{
		"Categories": categories,
	})
}
