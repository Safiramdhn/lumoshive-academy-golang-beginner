package main

import (
	"context"
	"fmt"
	"main/utils"
	"time"
)

type Data struct {
	role string
	nama string
}

func display(ctx context.Context, key string) {
	fmt.Println("Value Context :", ctx.Value(key))
}

func (d Data) addData(SliceData []Data, role string, name string, ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Akses ditolak: waktu habis!")
		return
	default:
		for _, data := range SliceData {
			if data.role == role && data.nama == name {
				fmt.Printf("Selamat datang %s (%s)\n", data.nama, data.role)
				ctxWithValue := context.WithValue(ctx, data.role, data.nama)
				display(ctxWithValue, data.role)
				showMenu(data.role)
				return
			}
		}
		fmt.Println("Data tidak valid!")
	}
}

func showMenu(role string) {
	switch role {
	case "admin":
		fmt.Println("==== Menu Admin ====")
		adminMenu := []string{"1. Dashboard", "2. User Management", "3. Reports"}

		for _, item := range adminMenu {
			fmt.Println(utils.ColorMessage("green", item))
		}
		fmt.Println("\n ")

	case "client":
		fmt.Println("==== Menu Client ====")
		clientMenu := []string{"1. Profile", "2. Orders", "3. Support"}
		for _, item := range clientMenu {
			fmt.Println(utils.ColorMessage("green", item))
		}
		fmt.Println("\n ")

	default:
		fmt.Println("Menu tidak ditemukan!")
	}
}

func main_latihan() {
	var SliceData []Data
	data := Data{}
	SliceData = append(SliceData, Data{role: "admin", nama: "haidar"})
	SliceData = append(SliceData, Data{role: "client", nama: "haidar2"})

	fmt.Printf("data : %v\n\n", SliceData)
	var role string
	var name string

	ctx := context.Background()
	deadline := time.Now().Add(10 * time.Second)
	ctxWithDeadline, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	for {
		fmt.Print("Masukan Role : ")
		fmt.Scan(&role)
		fmt.Print("Masukan Nama : ")
		fmt.Scan(&name)
		utils.ClearScreen()

		data.addData(SliceData, role, name, ctxWithDeadline)
		var choice string
		fmt.Print("Apakah Anda ingin mencoba lagi? (y/n): ")
		fmt.Scan(&choice)
		if choice != "y" {
			break
		}
	}
}
