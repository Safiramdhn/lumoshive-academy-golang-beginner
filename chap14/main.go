/*
	buatkan sebuah aplikasi reservasi restoran online yang di mana memiliki beberapa fitur
	- melakukan pesanan makanan
	- menambahkan pesanan makanan (edit)
	- melakukan pembayaran makanan
	- melihat pesanan yang sudah dibuat (history)
	- mengedit status
	- memberikan status pesanan ( di proses , di antar , selesai)
*/

package main

import (
	"fmt"
	"golang-beginner-14/controller"
	"os"
	"time"
)

func main() {
	for {
		var option int

		controller.ClearScreen()
		option = controller.PrintMainMenu()

		switch option {
		case 1:
			controller.CreateOrder()
		case 2:
			controller.EditOrder()
		case 3:
			controller.PayOrder()
		case 4:
			controller.PrintOrderHistory()
		case 5:
			controller.EditOrderStatus()
		case 99:
			os.Exit(0)
		default:
			fmt.Println("Invalid Input")
			time.Sleep(3 * time.Second)
		}
		continue
	}
}
