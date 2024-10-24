package controller

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
	fmt.Println("--------- Restaurant Reservation ---------")
	fmt.Println("1. Order Food")
	fmt.Println("2. Add Food (Edit)")
	fmt.Println("3. Payment")
	fmt.Println("4. Order History")
	fmt.Println("5. Edit Order Status")
	fmt.Println("Choose Option")
	fmt.Scan(&option)
	return option
}

func PromptReturnToMainMenu() {
	fmt.Println("---------")
	fmt.Println("Back to main menu? (y/n)")
	var response string
	fmt.Scan(&response)
	if response != "y" {
		os.Exit(0)
	}
	ClearScreen()
}

func promptContinue(action string) bool {
	fmt.Println("---------")
	fmt.Printf("%s/back to main menu? (y/n)\n", action)
	var response string
	fmt.Scan(&response)
	return response == "y"
}
