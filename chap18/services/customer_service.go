package services

import (
	"database/sql"
	"fmt"
	"golang-beginner-18/repositories"
)

func CountCustomerLogin(db *sql.DB) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction: %v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	customerDb := repositories.NewCustomerRepositoryDB(tx)
	totalCustomerLogin, totalCustomerLogout, err := customerDb.CountCustomerLogin()
	if err != nil {
		fmt.Printf("failed to get customer login/logout count: %v\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	fmt.Printf("Total Customer Login: %d\n", totalCustomerLogin)
	fmt.Printf("Total Customer Logout: %d\n", totalCustomerLogout)
	return nil
}

func GetFrequentCustomersByMonth(db *sql.DB) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction: %v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	customerDb := repositories.NewCustomerRepositoryDB(tx)
	customerOrderSumarries, err := customerDb.GetFrequentCustomersByMonth()
	if err != nil {
		fmt.Printf("failed to get frequent customers by month: %v\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	for _, orderSummary := range *customerOrderSumarries {
		// fmt.Printf("Month: %v, ID: %d, Name: %s, User Type: Customer, Total Orders: %d\n", orderSummary.Month, orderSummary.Id, orderSummary.Name, orderSummary.TotalOrders)
		fmt.Printf("ID: %d, Name: %s, User Type: Customer, Total Orders: %d\n", orderSummary.Id, orderSummary.Name, orderSummary.TotalOrders)

	}

	return nil
}

// func GetAllCusstomers(db *sql.DB) error {
// 	// Start a transaction
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Printf("failed to begin transaction: %v\n", err)
// 		return err
// 	}
// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	customerDb := repositories.NewCustomerRepositoryDB(tx)
// 	customers, err := customerDb.GetAll()
// 	if err != nil {
// 		fmt.Printf("failed to get all customers: %v\n", err)
// 		return err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		fmt.Printf("failed to commit transaction: %v\n", err)
// 		return err
// 	}

// 	for _, customer := range *customers {
// 		fmt.Printf("Customer ID: %d, First Name: %s, Last Name: %s, User ID: %d\n", customer.Id, customer.FirstName, customer.LastName, customer.UserId)
// 	}

// 	return nil
// }

// func GetCustomerById(db *sql.DB, id int) error {
// 	if id == 0 {
// 		fmt.Println("Invalid customer ID")
// 		return nil
// 	}

// 	// Start a transaction
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Printf("failed to begin transaction: %v\n", err)
// 		return err
// 	}
// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	customerDb := repositories.NewCustomerRepositoryDB(tx)
// 	customer, err := customerDb.GetById(id)
// 	if err != nil {
// 		fmt.Printf("failed to get customer by ID: %d, %v\n", id, err)
// 		return err
// 	}
// 	if err := tx.Commit(); err != nil {
// 		fmt.Printf("failed to commit transaction: %v\n", err)
// 		return err
// 	}

// 	if customer == nil {
// 		fmt.Printf("Customer not found with ID: %d\n", id)
// 		return nil
// 	}

// 	fmt.Printf("Customer ID: %d, First Name: %s, Last Name: %s, User ID: %d\n", customer.Id, customer.FirstName, customer.LastName, customer.UserId)
// 	return nil
// }
