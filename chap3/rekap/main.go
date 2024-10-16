package main

import "fmt"

func main() {
	dataProduct := []string{"kappa", "puma"}
	total := totalPrice(dataProduct)

	fmt.Println("Total yang harus dibayarkan Rp", total)
}

func totalPrice(merk []string) int {
	priceOfMerk := 0
	nominalDiscount := 0
	totalPayment := 0
	for _, value := range merk {
		if value == "adidas" {
			priceOfMerk += 200000
		} else if value == "puma" {
			priceOfMerk += 150000
		} else {
			priceOfMerk += 600000
		}
	}

	if len(merk) == 3 {
		nominalDiscount = 0
	} else {
		if checkDiscount(merk, "adidas") && checkDiscount(merk, "puma") {
			nominalDiscount = 50000
		} else if checkDiscount(merk, "puma") && checkDiscount(merk, "kappa") {
			nominalDiscount = 150000
		} else if checkDiscount(merk, "adidas") && checkDiscount(merk, "kappa") {
			nominalDiscount = 75000
		}
	}

	totalPayment = priceOfMerk - nominalDiscount
	return totalPayment
}

func checkDiscount(merk []string, selectMerk string) bool {
	for _, value := range merk {
		if value == selectMerk {
			return true
		}
	}
	return false
}
