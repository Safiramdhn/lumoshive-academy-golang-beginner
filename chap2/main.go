package main

import (
	"fmt"
)

func findName(target string, names []string) int {
	var result int
	for index, value := range names {
		if value == target {
			result = index
		}
	}
	return result
}

func setGrade(nilai int) string {
	var grade string
	switch {
	case nilai < 50:
		grade = "E"
	case nilai < 60:
		grade = "D"
	case nilai < 70:
		grade = "C"
	case nilai < 80:
		grade = "B"
	default:
		grade = "A"
	}
	return grade
}

func persegi(s int) (int, int) {
	var luas, keliling int
	luas = s * s
	keliling = 4 * s
	return luas, keliling
}

func discountCalculation(disc int, prices []int) []float64 {
	var total []float64
	for _, value := range prices {
		priceAfterDiscount := float64(value * (1 - (disc / 100)))
		total = append(total, priceAfterDiscount)
	}
	return total
}

func discountCondition(jml int, price int) (result float64) {
	if jml == 1 {
		result = float64(jml * price)
	} else if jml == 2 {
		result = float64(jml * (price - (price * 10 / 100)))
	} else if jml == 4 {
		result = float64(jml * (price - (price * 50 / 100)))
	} else if jml > 4 {
		result = float64(jml * (price - (price * 75 / 100)))
	}
	return result
}

func hitungGaji(kerja int, lembur int) (result int) {
	gajiKerjaPerjam := 50000
	gajiLemburPerjam := 60000

	totalGajiKerja := kerja * gajiKerjaPerjam
	totalGajiLembur := lembur * gajiLemburPerjam

	result = totalGajiKerja + totalGajiLembur
	return result
}

func hargaSepatu(merk []string) (total int) {
	hargaPuma := 150000
	hargaAdidas := 200000
	hargaKappa := 600000
	for i := 0; i < len(merk); i++ {
		// jika membeli sepatu adidas dan puma potongan 50.000
		if (merk[i] == "adidas" && merk[i+1] == "puma") || (merk[i] == "puma" && merk[i+1] == "adidas") {
			total = hargaAdidas + hargaPuma - 50000
			break

			// jika membeli sepatu puma dan kappa potongan 150.000
		} else if (merk[i] == "puma" && merk[i+1] == "kappa") || (merk[i] == "kappa" && merk[i+1] == "puma") {
			total = hargaPuma + hargaKappa - 150000
			break

			// jika membeli sepatu adidas dan kappa potongan 75.000
		} else if (merk[i] == "adidas" && merk[i+1] == "kappa") || merk[i] == "kappa" && merk[i+1] == "adidas" {
			total = hargaAdidas + hargaKappa - 75000
			break

			// selain kondisi diatas tidak dapat diskon
		} else {
			total = hargaAdidas + hargaKappa + hargaPuma
			break
		}
	}
	return total
}

func main() {
	// mencari index dari nama tertentu dari slice 5 data
	var names = []string{"safira", "bagas", "andi", "melina", "rara"}
	target := "safira"
	fmt.Printf("Nama %s di index %d\n", target, findName(target, names))

	// set grade
	nilai := 75
	fmt.Printf("Grade %d adalah %s\n", nilai, setGrade(nilai))

	// hitung luas dan keliling persegi
	sisiPersegi := 4
	luas, keliling := persegi(sisiPersegi)
	fmt.Printf("Luas persegi dengan sisi %d adalah %d\n", sisiPersegi, luas)
	fmt.Printf("keliling persegi dengan sisi %d adalah %d\n", sisiPersegi, keliling)

	// menampilkan harga setelah diskon yang ditentukan
	prices := []int{100000, 250000, 500000}
	priceAfterDiscount := discountCalculation(15, prices)
	fmt.Printf("Harga setelah discount %v\n", priceAfterDiscount)

	// menampilkan harga setelah diskon tergantung dengan jumlah pembelian
	jml := 5
	price := 20000
	harga := discountCondition(jml, price)
	fmt.Printf("Harga setelah diskon %.2f\n", harga)

	//function untuk hitung total gaji seorang karyawan jika 1 jam = 50rb,kemudian lembur 1 jam = 60rb. total kerja 40 jam, lembur 5 jam
	kerja := 40
	lembur := 5
	fmt.Printf("Total gaji seorang karyawa kerja %d jam dan lembur %d jam adalah %d\n", kerja, lembur, hitungGaji(kerja, lembur))

	/*
		buatkan function untuk studi kasus berikut :
		- diketahui ada product sepatu  adidas , puma, kappa
		- harga sepatu adidas 200.000
		- harga sepatu puma 150.000
		- harga sepatu kappa 600.000



		- harga diskon
		- jika membeli sepatu adidas dan puma potongan 50.000
		- jika membeli sepatu puma dan kappa potongan 150.000
		- jika membeli sepatu adidas dan kappa potongan 75.000
		- selain kondisi diatas tidak dapat diskon
	*/
	merk := []string{"adidas", "puma"}
	fmt.Printf("Total harga: %d\n", hargaSepatu(merk))
}
