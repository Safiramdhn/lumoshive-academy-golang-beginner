/*
	2/ buatkan 3 func yg fungsinya
	func 1 mencetak sebuah text setiap 2 detik
	func 2 mencetak text setiap 1 detik
	func 3 mencetak text setiap 3 detik
	buatkan context untuk membatalkan semua func yg berjalan di detik ke 5
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func printHelloWorld(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("Batalkan func 1\n")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("Hello, World!")
		}
	}
}

func printName(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("Batalkan func 2\n")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Hello, Safira!")
		}
	}
}

func printToday(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("Batalkan func 3\n")
			return
		default:
			time.Sleep(3 * time.Second)
			today := time.Now().Weekday()
			fmt.Printf("Today is %v\n", today)
		}
	}
}

func main() {
	parentCtx2 := context.Background()
	ctx3, cancel := context.WithCancel(parentCtx2)

	defer cancel()

	go printHelloWorld(ctx3)
	go printName(ctx3)
	go printToday(ctx3)

	time.Sleep(12 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")

}
