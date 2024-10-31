package repositories

import (
	"database/sql"
	"fmt"
)

type AdminRepositoryDB struct {
	DB *sql.DB
}

func NewAdminRepositoryDB(db *sql.DB) *AdminRepositoryDB {
	return &AdminRepositoryDB{DB: db}
}

func (repo *AdminRepositoryDB) Login(email, password string) (int, error) {
	sqlStatement := `
        SELECT a.id FROM users u
		JOIN admins a ON u.id = a.user_id
        WHERE u.email = $1 AND u.password = $2
    `
	var adminId int
	err := repo.DB.QueryRow(sqlStatement, email, password).Scan(&adminId)

	fmt.Printf("err: %v\n", err)

	if err == sql.ErrNoRows {
		fmt.Println("No user found with the given email and password.")
		return 0, nil // User not found
	} else if err != nil {
		fmt.Printf("QueryRow error: %v\n", err)
		return 0, err // Error executing query
	}

	return adminId, nil
}
