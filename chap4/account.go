package main

import "errors"

type Saldo struct {
	Saldo float64
}
type User struct {
	Name  string
	Email string
	Saldo Saldo
}

type Accounts struct {
	accounts []User
}

type AccountManager interface {
	createAccount(name, email string) (User, error)
	getAccount() []User
}

type SaldoManager interface {
	addBalance(amount float64) (User, error)
	reduceBalance(amount float64) (User, error)
}

func (a *Accounts) createAccount(name, email string) (User, error) {
	if name == "" || email == "" {
		return User{}, errors.New("Nama dan Email tidak boleh kosong")
	}

	for _, account := range a.accounts {
		if account.Email == email {
			return User{}, errors.New("Email sudah digunakan")
		}
	}

	newUser := User{Name: name, Email: email, Saldo: Saldo{Saldo: 0.0}}
	a.accounts = append(a.accounts, newUser)
	return newUser, nil
}

func (a *Accounts) getAccount() []User {
	return a.accounts
}

func (a *Accounts) addBalance(email string, amount float64) (User, error) {
	var updatedUser User
	count := 0
	for i, account := range a.accounts {
		count++
		if account.Email == email {
			account.Saldo.Saldo += amount
			updatedUser = account
			a.accounts = append(a.accounts[:i], updatedUser)
		} else {
			if count == len(a.accounts) {
				return User{}, errors.New("User tidak ditemukan")
			}
		}
	}
	return updatedUser, nil
}

func (a *Accounts) reduceBalance(email string, amount float64) (User, error) {
	var updatedUser User
	count := 0
	for i, account := range a.accounts {
		if account.Email == email {
			account.Saldo.Saldo -= amount
			updatedUser = account
			a.accounts = append(a.accounts[:i], updatedUser)
		} else {
			if count == len(a.accounts) {
				return User{}, errors.New("User tidak ditemukan")
			}
		}
	}
	return updatedUser, nil
}
