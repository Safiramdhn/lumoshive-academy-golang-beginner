package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func PrintMainMenu() (option int) {
	fmt.Println("--------- Main Menu ---------")
	fmt.Println("1. Product List")
	fmt.Println("2. Add Product")
	fmt.Println("3. Cart")
	fmt.Println("4. Checkout")
	fmt.Println("5. Log out")
	fmt.Println("Choose Option :")
	fmt.Scan(&option)
	return option
}
