/*
buatkan sistem keuangan untuk mengelola data keuangan seseorang
	sistem dapat menambahkan saldo
	sistem dapat mengurangi saldo
	sistem dapat membuat akun baru
	sistem dapat mencetak akun berhasil di tambahkan, akun berhasil menambahkan saldo, akun berhasil mengurangi saldo

	ketentuan baku
	sistem memiliki 2 object account dan saldo
	buatkan 2 object tersebut di file terpisah di satu package
	buatkan sistem validasi error saat manambahkan akun , menambah saldo dan mengurangi saldo

	contoh output :
	akun berhasil di tambahkan [{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } 3]
	saldt berhasil di tambahkan [{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } 3]
	saldo berhasil di dikurangi [{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 3]
*/

package main

import (
	"fmt"
)

func main() {
	Accounts := Accounts{}
	account1, err := Accounts.createAccount("lumo", "lumoshive@gmail.com")
	if err != nil {
		fmt.Println("Error Message:", err)
	} else {
		fmt.Println("Account berhasil di tambahkan", account1)
	}

	account2, err := Accounts.createAccount("safira", "safira9@gmail.com")
	if err != nil {
		fmt.Println("Error Message:", err)
	} else {
		fmt.Println("Account berhasil di tambahkan", account2)
	}

	accountList := Accounts.getAccount()
	fmt.Printf("Account List : \n%+v\n", accountList)

	updatedUser, err := Accounts.addBalance("safira9@gmail.com", 50000.0)
	if err != nil {
		fmt.Println("Error Message:", err)
	} else {
		fmt.Println("Saldo berhasil ditambahkan", updatedUser)
	}

	updatedUser2, err := Accounts.reduceBalance("safira9@gmail.com", 5000.0)
	if err != nil {
		fmt.Println("Error Message:", err)
	} else {
		fmt.Println("Saldo berhasil dikurangkan", updatedUser2)
	}

}
