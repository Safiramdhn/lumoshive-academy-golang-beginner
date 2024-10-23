/*
	buat slice dengan 2 data struct
	struct role dan nama
	role memiliki akses menu yang berbeda
	buat 1 func berufungsi untuk cek 2 data struct dalam slice
	kondisi :
	- if ada masukan ke context value

	buat func untuk menampilkan beberapa menu sesuai dengan data yang dipil
	pengecekan data tida bisa diakses dalam waktu tertertentu with deadline
*/

package main

import (
	"context"
	"fmt"
	"time"
)

type Role struct {
	Name  string
	Menus []string
}

type User struct {
	Name string
	Role Role
}

var users = []User{
	{Name: "Alice", Role: Role{Name: "Admin", Menus: []string{"Dashboard", "Settings", "Users", "Profile"}}},
	{Name: "Bob", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Caterina", Role: Role{Name: "Admin ", Menus: []string{"Dashboard", "Settings", "Users"}}},
	{Name: "Dan", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Eline", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Farah", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Grey", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Hugh", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Irene", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Profile"}}},
	{Name: "Jane", Role: Role{Name: "User ", Menus: []string{"Dashboard", "Users"}}},
}

func checkAccess(user User, menu string) bool {
	for _, m := range user.Role.Menus {
		if m == menu {
			return true
		}
	}
	return false
}

func displayMenu(ctx context.Context, user User) {
	select {
	case <-ctx.Done():
		fmt.Println("Access denied: deadline has passed.")
		return
	default:
		fmt.Printf("Menu for %s (%s):\n", user.Name, user.Role.Name)
		for _, menu := range user.Role.Menus {
			fmt.Println("-", menu)
		}
	}
}

func main() {
	// Atur deadline 1 jam dari sekarang
	deadline := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // Pastikan untuk membatalkan konteks setelah selesai

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Access denied: deadline has passed.")
			return
		default:
			for _, user := range users {
				time.Sleep(4 * time.Second)
				// Cek apakah konteks telah dibatalkan
				if checkAccess(user, "Settings") {
					displayMenu(ctx, user)
				} else {
					fmt.Printf("%s does not have access to Settings.\n", user.Name)
				}
			}
		}
	}
}
