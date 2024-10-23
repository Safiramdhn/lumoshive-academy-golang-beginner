package product

import (
	"fmt"
	"golang-beginner-13/task/utils"
	"time"
)

type Product struct {
	Id    int
	Name  string
	Price float64
}

type Cart struct {
	Product []Product
}

func AddProduct(product_id int, c Cart) Cart {
	productList := productList()
	productFound := false
	for _, product := range productList {
		if product_id == product.Id {
			fmt.Println("Product added to cart.")
			c.Product = append(c.Product, product)
			productFound = true
			break
		}
	}

	if !productFound {
		fmt.Println("Product not found")
	}

	return c
}

func ShowCart(c Cart) {
	for i, product := range c.Product {
		fmt.Printf("%d Name: %s, Price: Rp %.2f\n", i+1, product.Name, product.Price)
	}
}

func Checkout(c Cart) Cart {
	for {

		ShowCart(c)
		var pay float64
		totalPrice := 0.0
		totalItem := 0
		for _, product := range c.Product {
			totalPrice += product.Price
			totalItem++
		}

		fmt.Printf("Total Item : %d\n", totalItem)
		fmt.Printf("Total Price : Rp %.2f\n", totalPrice)
		fmt.Println("Pay (Rp): ")
		fmt.Scan(&pay)

		if pay < totalPrice {
			fmt.Printf("You must pay as much as %.2f", totalPrice)
			time.Sleep(3 * time.Second)
			utils.ClearScreen()
			continue
		}
		fmt.Printf("Change : Rp %.2f\n", pay-totalPrice)
		fmt.Println("Goods will be shipped soon")
		c.Product = []Product{}
		return c
	}
}

func productList() []Product {
	productList := []Product{
		{1, "Apple", 10000.0},
		{2, "Lemon", 12000.0},
		{3, "Melon", 20000.0},
		{4, "Strawberry", 30000.0},
		{5, "Blueberry", 35000.0},
		{6, "Dragon Fruit", 35000.0},
		{7, "Kiwi", 32000.0},
		{8, "Watermelon", 35000.0},
		{9, "Orange", 15000.0},
		{10, "Lychee", 21000.0},
	}
	return productList
}

func PrintProductList() {
	fmt.Println("--------- Product List ---------")
	productList := productList()
	for _, list := range productList {
		fmt.Printf("ID : %d\t| Name: %s\t | Price: %.2f\t|\n", list.Id, list.Name, list.Price)
	}
}
