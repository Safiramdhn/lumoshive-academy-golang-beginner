package main

import (
	"fmt"
	"reflect"
)

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
	kecepatan int
}

type Motor struct {
	kecepatan int
}

type Bajaj struct {
	kecepatan int
}

type Kendaraan interface {
	Efesiensi(gas int) int
}

func (mobil Mobil) Efesiensi(gas int) int {
	return mobil.kecepatan * gas
}

func (motor Motor) Efesiensi(gas int) int {
	return motor.kecepatan * gas
}

func (bajaj Bajaj) Efesiensi(gas int) int {
	return bajaj.kecepatan * gas
}

func CalculateEfesiensi(kendaraan []Kendaraan, gas int) string {
	var efesien Kendaraan
	max := 0

	for _, item := range kendaraan {
		jarak := item.Efesiensi(gas)
		if jarak > max {
			max = jarak
			efesien = item
		}
	}

	switch reflect.TypeOf(efesien).Name() {
	case "Mobil":
		return "Mobil lebih efesien"
	case "Motor":
		return "Motor lebih efesien"
	case "Bajaj":
		return "Bajaj lebih efesien"
	default:
		return "Tidak ada kendaraan terdaftar"
	}

}

func main() {
	mobil := Mobil{kecepatan: 1}
	motor := Motor{kecepatan: 3}
	bajaj := Bajaj{kecepatan: 4}

	listKendaraan := []Kendaraan{mobil, motor, bajaj}
	result := CalculateEfesiensi(listKendaraan, 10)
	fmt.Println(result)
}
