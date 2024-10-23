/*
	buatkan 1 variable tipe int, inisialisasikan dengan jumlah tertentu.
	buat sebuah fungsi untuk mengurangi nilai dari variable itu, fungsi itu dijalankan pakai goroutine, pengurangan dilakukan setiap 2 detik.
	berikan context timeout pada detik ke 4.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	num := 99
	parentCtx3 := context.Background()
	ctx4, cancel2 := context.WithTimeout(parentCtx3, 4*time.Second)
	defer cancel2()

	go reduceNum(ctx4, num)
	<-ctx4.Done()
}

func reduceNum(ctx context.Context, num int) {
	for i := 0; i < num; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Pengurangan berhenti pada angka %d\n", num)
			return
		default:
			if num != 0 {
				time.Sleep(2 * time.Second)
				num--
				fmt.Printf("Current number: %d\n", num)
			} else {
				fmt.Println("Number has been reduced to zero")
			}
		}
	}
}
