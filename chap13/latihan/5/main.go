/*
3/ buatkan sebuah func untuk menambahkan nilai ke dalam context (key dan value)
	dan buatkan 1 func untuk menampilkan key value setiap 2 detik
	buat 1 func untuk mengubah nilai value sebelumnya

5/ buatkan 1 struct dgn 2 properti tipe data bebas
	buatkan 1 func untuk memasukan data struct ke dalam slice
	buatkan context menggunakan deadline 10 detik kedepan func berhenti
*/

package main

import (
	"context"
	"fmt"
	"time"
)

type Data struct {
	id           int
	time_created string
}

func main() {

	parentCtx := context.Background()
	ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(10*time.Second))
	defer cancel()
	count := 0
	storage := []Data{}

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Deadline exceeded. Total Data: %d\n", count)
			return
		default:
			now := time.Now().Format("01/02 03:04:05PM")
			data := Data{id: count + 1, time_created: now}
			time.Sleep(3 * time.Second)
			storage = insertDataSlice(data, storage)
			count++
		}
	}
}

func insertDataSlice(data Data, storage []Data) []Data {
	storage = append(storage, data)
	fmt.Printf("Storage Value Bertambah %+v\n", storage)
	return storage
}
