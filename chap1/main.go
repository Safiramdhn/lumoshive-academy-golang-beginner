package main

import (
	"fmt"
)

func main() {
	/*
		buat variable slice jml data 20 (random)
		print data index 8 s/d akhir
	*/
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(numbers[8:])

	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	diskon := 15.0
	jumlahTerjual := 100

	hargaJualSetelahDiskon := hargaJual - (hargaJual * diskon / 100.0)
	totalPendapatan := int(hargaJualSetelahDiskon) * jumlahTerjual
	totalBiaya := hargaBeli + biayaOperasional
	totalKeuntungan := int(totalPendapatan) - int(totalBiaya*float64(jumlahTerjual))

	fmt.Println("Harga Jual Setelah Diskon: ", hargaJualSetelahDiskon)
	fmt.Println("Total Pendapatan: ", totalPendapatan)
	fmt.Println("Total Biaya: ", totalBiaya)
	fmt.Println("Total Keuntungan: ", totalKeuntungan)

	data := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(data[2][2])
}
