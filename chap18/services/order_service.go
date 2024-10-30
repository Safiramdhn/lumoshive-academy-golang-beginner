package services

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-beginner-18/models"
	"golang-beginner-18/repositories"
	"log"
	"time"
)

func CreateOrder(db *sql.DB, orderInput *models.Orders) error {
	// Check for required fields
	if orderInput.CustomerId == 0 || orderInput.DriverId == 0 {
		return errors.New("customer_id and driver_id must be provided")
	}

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Defer rollback for safety, only commit if successful
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	// Create a repository instance
	orderDB := repositories.NewOrdersRepositoryDB(tx)

	// Format date and time for consistent insertion
	orderInput.OrderDate = orderInput.OrderDate.Truncate(24 * time.Hour) // Ensure date-only format
	orderInput.OrderTime = orderInput.OrderTime.Truncate(time.Second)    // Ensure time-only format

	// Insert order using the repository
	orderID, err := orderDB.Create(orderInput)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	fmt.Printf("Order created successfully with ID: %d\n", orderID)
	return nil
}

func GetTotalOrder(db *sql.DB) error {
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

	orderDB := repositories.NewOrdersRepositoryDB(tx)
	totalOrders, err := orderDB.GetTotalOrder()
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	for _, order := range *totalOrders {
		fmt.Printf("Month: %s, Total Orders: %d\n", order.Month, order.TotalOrders)
	}
	return nil
}

func GetPopularAreas(db *sql.DB) error {
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

	orderDB := repositories.NewOrdersRepositoryDB(tx)
	popularAreas, err := orderDB.GetPopularAreas()
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	for _, area := range *popularAreas {
		fmt.Printf("City: %s, District: %s, Neighborhood: %s, Street Name: %s, Total Orders: %d\n", area.City, area.District, area.Neighborhood, area.StreetName, area.TotalOrders)
	}
	return nil
}

func GetOrderPeakHours(db *sql.DB) error {
	// Begin a new transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("failed to begin transaction: %v\n", err)
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

	// Initialize repository and get peak hours
	orderDB := repositories.NewOrdersRepositoryDB(tx)
	orderPeakHours, err := orderDB.GetOrderPeakHours()
	if err != nil {
		tx.Rollback() // Rollback on error
		return err
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	// Display the results
	for _, peakHour := range orderPeakHours {
		fmt.Printf("Hour Range: %s, Total Orders: %d\n", peakHour.HourRange, peakHour.TotalOrders)
	}

	return nil
}
