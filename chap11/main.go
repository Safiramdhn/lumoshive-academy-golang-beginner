/*
- buatkan sebuah program untuk menangani permintaan penambahan data ke slice sebanyak 100 data
- gunakan go routine untuk menangani func untuk proses penyimpanan data ke slice
- gunakan go routine untuk mencetak di log data berhasil di simpan di slice
- buatkan satu struct untuk model datanya yang disimpan di struct
- terapkan juga pattern code yang modular.
*/

package main

import (
	"fmt"
	"golang-beginner-11/model"
	"golang-beginner-11/service"
)

func main() {
	// Buatkan channel untuk komunikasi antara goroutine
	dataChan := make(chan model.Data)
	doneChan := make(chan bool)

	// Buatkan slice untuk menyimpan data
	var dataSlice []model.Data

	// Jalankan goroutine untuk proses penyimpanan data
	go service.ProcessData(dataChan, &dataSlice)

	// Proses penambahan data ke channel
	for i := 1; i <= 100; i++ {
		go func(id int) {
			fmt.Println("Mengirim data ke-", i)
			dataChan <- model.Data{Id: id, Name: fmt.Sprintf("Data %d", id)}
			doneChan <- true
		}(i)
	}

	// Tunggu sampai pengguna menekan enter
	fmt.Println("Tekan enter untuk melihat hasil")
	fmt.Scanln()

	// Tutup channel
	<-doneChan

	// Tunggu sampai semua goroutine selesai
	fmt.Printf("Total data disimpan: %d\n", len(dataSlice))
}
