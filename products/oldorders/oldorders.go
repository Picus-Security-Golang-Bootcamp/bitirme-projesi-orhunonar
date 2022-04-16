package oldorders

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OldOrders struct {
	gorm.Model
	Name  string  `csv:"name"`
	Price float64 `csv:"price"`
	Stock int     `csv:"stock"`
}

func AddOldOrder(Name string, Stock int) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&OldOrders{})
	if err != nil {
		panic(err)
	}
	oldorders := OldOrders{Name: Name, Stock: Stock}
	db.Create(&oldorders)

}

func RemoveOldOrder(Name string) {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&OldOrders{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).Delete(&OldOrders{})

}
func GetOldOrders(c *gin.Context) []OldOrders {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&OldOrders{})
	if err != nil {
		panic(err)
	}
	var oldorders []OldOrders
	db.Find(&oldorders)
	c.JSON(200, gin.H{
		"oldorders": oldorders,
	})
	return oldorders

}
