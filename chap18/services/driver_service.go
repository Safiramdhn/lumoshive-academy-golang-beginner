package services

import (
	"database/sql"
	"fmt"
	"golang-beginner-18/repositories"
)

func CountDriverLogin(db *sql.DB) error {
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

	driverDb := repositories.NewDriverRepositoryDB(tx)
	totalDriverLogin, totalDriverLogout, err := driverDb.CountDriverLogin()
	if err != nil {
		fmt.Printf("failed to get driver login/logout count: %v\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	fmt.Printf("Total Driver Login: %d\n", totalDriverLogin)
	fmt.Printf("Total Driver Logout: %d\n", totalDriverLogout)
	return nil
}

func GetFrequentDriversByMonth(db *sql.DB) error {
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

	driverDb := repositories.NewDriverRepositoryDB(tx)
	customerOrderSumarries, err := driverDb.GetActiveDriversByMonth()
	if err != nil {
		fmt.Printf("failed to get active driver by month: %v\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	for _, orderSummary := range *customerOrderSumarries {
		// fmt.Printf("Month: %v, ID: %d, Name: %s, User Type: Driver, Total Orders: %d\n", orderSummary.Id, orderSummary.Name, orderSummary.TotalOrders)
		fmt.Printf("ID: %d, Name: %s, User Type: Driver, Total Orders: %d\n", orderSummary.Id, orderSummary.Name, orderSummary.TotalOrders)
	}

	return nil
}
