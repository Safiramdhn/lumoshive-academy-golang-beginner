package services

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-beginner-18/models"
	"golang-beginner-18/repositories"
)

func CreateUser(db *sql.DB, firstName, lastName, email, password, role string) error {
	// Validate inputs
	if firstName == "" || lastName == "" || email == "" || password == "" {
		return errors.New("all fields must be filled")
	}

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

	// Create user within transaction
	userDB := repositories.NewUserRepositoryDB(tx)
	userId, err := userDB.Create(email, password)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return err
	}

	switch role {
	case "customer":
		customerDB := repositories.NewCustomerRepositoryDB(tx)
		customer := models.Customer{UserId: userId, FirstName: firstName, LastName: lastName}
		if err = customerDB.Create(&customer); err != nil {
			fmt.Printf("failed to create customer: %v\n", err)
		}
	case "driver":
		driverDB := repositories.NewDriverRepositoryDB(tx)
		driver := models.Driver{Id: userId, FirstName: firstName, LastName: lastName}
		if err = driverDB.Create(&driver); err != nil {
			fmt.Printf("Error creating driver: %v\n", err)
			return err
		}
	default:
		return errors.New("invalid role specified")
	}

	// Commit the transaction if no errors occurred
	if err = tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}

	fmt.Println("User created successfully!")
	return nil
}

func LoginUser(db *sql.DB, email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password must be provided")
	}

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

	userDB := repositories.NewUserRepositoryDB(tx)
	_, err = userDB.Login(email, password)
	if err != nil {
		fmt.Printf("Error logging in user: %v\n", err)
		return err
	}
	// Commit the transaction if no errors occurred
	if err = tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}
	fmt.Println("User logged in successfully!")
	return nil
}

func LogoutUser(db *sql.DB, userId int) error {
	if userId == 0 {
		return errors.New("invalid user ID")
	}

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

	userDB := repositories.NewUserRepositoryDB(tx)
	_, err = userDB.Logout(userId)
	if err != nil {
		fmt.Printf("Error logging out user: %v\n", err)
		return err
	}
	// Commit the transaction if no errors occurred
	if err = tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}
	fmt.Println("User logged out successfully!")
	return nil
}

func DeleteUser(db *sql.DB, userId int) error {
	if userId == 0 {
		return errors.New("invalid user ID")
	}

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

	userDB := repositories.NewUserRepositoryDB(tx)
	err = userDB.Delete(userId)
	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		return err
	}
	// Commit the transaction if no errors occurred
	if err = tx.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %v\n", err)
		return err
	}
	fmt.Println("User deleted successfully!")
	return nil
}
