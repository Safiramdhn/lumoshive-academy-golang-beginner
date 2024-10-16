package main

import "fmt"

/*
Soal bersama :
- buatkan 3 object struct kendaraan (motor , mobil , bajaj)
- masing-masing object struct tersebut memiliki properttkecepatan
- buatkan function dari masing-masing struct dengan spesifikasi :
- 1 liter = untuk 1 km (mobil)
- 1 hter = untuk 3 km (motor)
- 1 liter = untuk 4 km (bajaj)
- buatkan satu function untuk menentukan mana yang paling efesien dari 3 kendaraan tersebut,
function memiliki paramter bensi dan object kendara (bisa langsung di masukkan 3 object)
- jika memiliki 10 liter bensin
*/

type Mobil struct {
	kecepatan       int
	jarak_per_liter int
}

type Motor struct {
	kecepatan       int
	jarak_per_liter int
}

type Bajaj struct {
	kecepatan       int
	jarak_per_liter int
}

type Kendaraan struct {
	Mobil
	Motor
	Bajaj
}

type Efesiensi interface {
	Efesiensi(gas int) int
}

func (k Kendaraan) Efesiensi(gas int) string {
	k.Mobil.kecepatan = k.Mobil.jarak_per_liter * gas
	k.Motor.kecepatan = k.Motor.jarak_per_liter * gas
	k.Bajaj.kecepatan = k.Bajaj.jarak_per_liter * gas

	if (k.Mobil.kecepatan > k.Motor.kecepatan) && (k.Mobil.kecepatan > k.Bajaj.kecepatan) {
		return "Mobil lebih efesien"
	} else if (k.Motor.kecepatan > k.Bajaj.kecepatan) && (k.Motor.kecepatan > k.Mobil.kecepatan) {
		return "Motor lebih efesien"
	} else {
		return "Bajaj lebih efesien"
	}
}

func CalculateEfesiensi(kendaraan Kendaraan, gas int) {
	fmt.Println(kendaraan.Efesiensi(gas))
}

func main() {
	kendaraan := Kendaraan{Mobil{jarak_per_liter: 5}, Motor{jarak_per_liter: 3}, Bajaj{jarak_per_liter: 4}}
	CalculateEfesiensi(kendaraan, 25)
}
