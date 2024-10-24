/*
Buatkan aplikasi belanja online yang terdiri dari beberapa menu :
- login (buatkan sesi login 20 detik) jika sesi habis maka saat memilih menu menampilkan pesan silakan
login kembali
- tampilkan beberapa product
- pilih product untuk di tambahkan ke keranjang
- lihat keranjang
- checkout product dan bayar sesuai nominal kemudian tampilkan pesan baranga segera di kirim

ketentuan
- terapkan context
- slice
- struct
- function
- scan data dari terminal
*/

package main

import (
	"context"
	"fmt"
	"golang-beginner-13/task/product"
	user "golang-beginner-13/task/users"
	"golang-beginner-13/task/utils"
	"os"
	"time"
)

func init() {
	utils.ClearScreen()
	fmt.Println("Welcome to online shopping apps")
}

func main() {
	person := user.Person{
		Username: "client",
		Password: "password",
	}
	users := user.Users{}
	users = user.AddUsers(person, users)
	var cart = product.Cart{} // The cart is initialized once and persists
	parentContext := context.Background()

loginLoop:
	for {
		// Use utils functions, no direct dependency on product in utils
		username, password := promptLogin()
		err := user.Login(username, password, users)
		if err != nil {
			fmt.Printf("Error message: %s\n", err)
			return
		}
		fmt.Println("Login Success")

		ctx, cancel := context.WithTimeout(parentContext, 20*time.Second)
		defer cancel()

		handleMainMenu(ctx, &cart) // Pass the cart by reference

		// When the session expires or logs out, continue loginLoop but the cart persists
		continue loginLoop
	}
}

func promptLogin() (string, string) {
	var username, password string
	fmt.Println("--------- Welcome ------------")
	fmt.Println("Username :")
	fmt.Scan(&username)
	fmt.Println("Password :")
	fmt.Scan(&password)
	return username, password
}

func handleMainMenu(ctx context.Context, cart *product.Cart) {
	for {
		utils.ClearScreen()
		select {
		case <-ctx.Done():
			fmt.Println("Session expired, please login again.")
			time.Sleep(5 * time.Second)
			utils.ClearScreen()
			// Go back to login loop, but keep the cart intact
			return
		default:
			option := utils.PrintMainMenu()
			switch option {
			case 1:
				handleProductList()
			case 2:
				handleAddToCart(cart) // Pass the existing cart
			case 3:
				handleViewCart(*cart) // Pass the cart by value for viewing
			case 4:
				handleCheckout(cart) // Pass the existing cart for checkout
			case 5:
				if confirmLogout() {
					fmt.Println("Thank you")
					time.Sleep(2 * time.Second)
					os.Exit(0)
				}
			default:
				fmt.Println("Invalid Option")
			}
		}
	}
}

func handleProductList() {
	utils.ClearScreen()
	product.PrintProductList()
	promptReturnToMainMenu()
}

func handleAddToCart(cart *product.Cart) {
	for {
		utils.ClearScreen()
		fmt.Println("--------- Add To Cart ---------")
		var id int
		product.PrintProductList()
		fmt.Println("Enter product id :")
		fmt.Scan(&id)
		*cart = product.AddProduct(id, *cart) // Calls the updated AddProduct function

		if !promptContinue("Add new item") {
			utils.ClearScreen()
			return
		}
	}
}

func handleViewCart(cart product.Cart) {
	utils.ClearScreen()
	fmt.Println("--------- Cart ---------")
	if len(cart.Product) == 0 {
		fmt.Println("No products added to the cart.")
		time.Sleep(3 * time.Second)
		return
	}
	product.ShowCart(cart)
	promptReturnToMainMenu()
}

func handleCheckout(cart *product.Cart) {
	fmt.Println("--------- Check ---------")
	utils.ClearScreen()
	if len(cart.Product) == 0 {
		fmt.Println("Please add products to the cart first.")
		time.Sleep(3 * time.Second)
		return
	}
	*cart = product.Checkout(*cart)
	promptReturnToMainMenu()
}

func confirmLogout() bool {
	fmt.Println("---------")
	fmt.Println("If you log out, your data will be removed. Continue? (y/n)")
	var response string
	fmt.Scan(&response)
	return response == "y"
}

func promptReturnToMainMenu() {
	fmt.Println("---------")
	fmt.Println("Back to main menu? (y/n)")
	var response string
	fmt.Scan(&response)
	if response != "y" {
		os.Exit(0)
	}
	utils.ClearScreen()
}

func promptContinue(action string) bool {
	fmt.Println("---------")
	fmt.Printf("%s/back to main menu? (y/n)\n", action)
	var response string
	fmt.Scan(&response)
	return response == "y"
}
