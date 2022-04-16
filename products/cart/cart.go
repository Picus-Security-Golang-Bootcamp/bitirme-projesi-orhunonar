package cart

import (
	"finalproject/products/categories"
	"finalproject/products/oldorders"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Name  string  `csv:"name"`
	Price float64 `csv:"price"`
	Stock int     `csv:"stock"`
}

//Add a shoe to the cart
func AddShoeToCart(Name string, Stock int) {

	var shoe categories.Shoes
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).First(&shoe)
	if shoe.Name == Name {
		if shoe.Stock >= Stock {

			shoesincart := Cart{Name: Name, Stock: Stock, Price: shoe.Price}
			db.Create(&shoesincart)

		} else {
			log.Fatal("Not enough stock")
		}

	} else {
		log.Fatal("Shoe not found")
	}

}

//Add a product to the cart
func AddPantsToCart(Name string, Stock int) {

	var pants categories.Pants
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).First(&pants)
	if pants.Name == Name {
		if pants.Stock >= Stock {

			pantsincart := Cart{Name: Name, Stock: Stock, Price: pants.Price}
			db.Create(&pantsincart)

		} else {
			log.Fatal("Not enough stock")
		}

	} else {
		log.Fatal("Pants not found")
	}

}

//Add a glasses to the cart
func AddGlassesToCart(Name string, Stock int) {

	var glasses categories.Glasses
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).First(&glasses)
	if glasses.Name == Name {
		if glasses.Stock >= Stock {

			glassesincart := Cart{Name: Name, Stock: Stock, Price: glasses.Price}
			db.Create(&glassesincart)

		} else {
			log.Fatal("Not enough stock")
		}

	} else {
		log.Fatal("Glasses not found")
	}

}

// RemoveProductFromCart removes a product from the cart
func RemoveProductFromCart(Name string) {

	var products []Cart
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).Find(&products)
	if products[0].Name == Name {
		db.Where("name = ?", Name).Delete(&products)
	} else {
		log.Fatal("Product not found")
	}

}

//Update the stock of a product
func UpdateProductStock(Name string, Stock int) {

	var products []Cart
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).Find(&products)
	if products[0].Name == Name {
		if products[0].Stock >= Stock {
			db.Model(&products).Update("stock", Stock)
		} else {
			log.Fatal("Not enough stock")
		}
	} else {
		log.Fatal("Product not found")
	}

}

//List all products in cart
func ListProducts(c *gin.Context) {
	var products []Cart

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Find(&products)

	for _, product := range products {
		log.Println(product.Name)
	}
	c.JSON(200, products)

}

//Get the total price of the cart
func GetTotalPrice(c *gin.Context) float64 {
	var products []Cart
	var totalprice float64

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Find(&products)

	for _, product := range products {
		totalprice += product.Price
	}
	c.JSON(200, totalprice)

	return totalprice

}
func CompleteOrder() {

	var products []Cart
	var shoe categories.Shoes
	var pants categories.Pants
	var glasses categories.Glasses
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})

	if err != nil {
		panic(err)
	}
	db.Find(&products)

	for _, product := range products {
		if product.Name == shoe.Name {
			db.Where("name = ?", product.Name).First(&shoe)
			db.Model(&shoe).Update("stock", shoe.Stock-product.Stock)
		} else if product.Name == pants.Name {
			db.Where("name = ?", product.Name).First(&pants)
			db.Model(&pants).Update("stock", pants.Stock-product.Stock)
		} else if product.Name == glasses.Name {
			db.Where("name = ?", product.Name).First(&glasses)
			db.Model(&glasses).Update("stock", glasses.Stock-product.Stock)
		}
		oldorders.AddOldOrder(product.Name, product.Stock)
	}
	db.Delete(&products)

}

func ClearCart() {
	var products []Cart
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Find(&products)

	for _, product := range products {
		db.Where("name = ?", product.Name).Delete(&product)
	}

}

func CancelOrder(c *gin.Context) {

	//if 14 days has not been passed, the order can be cancelled
	if oldorders.CheckDate(c) {
		var products []Cart
		var shoe categories.Shoes
		var pants categories.Pants
		var glasses categories.Glasses
		db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		err = db.AutoMigrate(&Cart{})
		if err != nil {
			panic(err)
		}
		db.Find(&products)

		for _, product := range products {
			if product.Name == shoe.Name {
				db.Where("name = ?", product.Name).First(&shoe)
				db.Model(&shoe).Update("stock", shoe.Stock+product.Stock)
			} else if product.Name == pants.Name {
				db.Where("name = ?", product.Name).First(&pants)
				db.Model(&pants).Update("stock", pants.Stock+product.Stock)
			} else if product.Name == glasses.Name {
				db.Where("name = ?", product.Name).First(&glasses)
				db.Model(&glasses).Update("stock", glasses.Stock+product.Stock)
			}
		}
		db.Delete(&products)
	} else {
		log.Fatal("Order can not be cancelled")
	}

}
func SearchProductInCart(Name string, c *gin.Context) {
	var products []Cart
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=admin dbname=Bucket port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Cart{})
	if err != nil {
		panic(err)
	}
	db.Where("name = ?", Name).Find(&products)
	if products[0].Name == Name {
		c.JSON(200, products)
	} else {
		log.Fatal("Product not found")
	}

}

//Check if 14 days passed from today

type Date struct {
	gorm.Model
	Date string
}
