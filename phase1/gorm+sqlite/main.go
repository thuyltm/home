package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&Product{})
	//Create
	db.Create(&Product{Code: "D42", Price: 100})

	var products []Product
	result := db.Find(&products)

	if result.Error != nil {
		fmt.Println("Error finding products:", result.Error)
		return
	}

	// Print the results using fmt.Println
	fmt.Println("Found products:", products)

	//Read
	var product Product
	result = db.First(&product)
	if result.Error != nil {
		fmt.Println("Error finding products:", result.Error)
		return
	}

	// Print the results using fmt.Println
	fmt.Println("Found product:", product)

	db.First(&product, "code = ?", "D42")
	//Update
	db.Model(&product).Update("Price", 200)
	//Update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	// Delete
	db.Delete(&product, 1)
}
