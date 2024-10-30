package repositories

import (
	"database/sql"
)

type UserRepositoryDB struct {
	DB *sql.Tx
}

func NewUserRepositoryDB(db *sql.Tx) *UserRepositoryDB {
	return &UserRepositoryDB{DB: db}
}

func (repo *UserRepositoryDB) Create(email, password string) (int, error) {
	var userId int
	sqlStatement := `
        INSERT INTO users (email, password)
        VALUES ($1, $2) RETURNING id
    `
	err := repo.DB.QueryRow(sqlStatement, email, password).Scan(&userId)
	return userId, err
}

func (repo *UserRepositoryDB) Login(email, password string) (bool, error) {
	sqlStatement := `
        SELECT id FROM users
        WHERE email = $1 AND password = $2
    `
	var userId int
	err := repo.DB.QueryRow(sqlStatement, email, password).Scan(&userId)
	if userId > 0 {
		updateStatement := ` UPDATE users SET login_time = NOW() logout_time = NULL WHERE id = $1;`
		_, err = repo.DB.Exec(updateStatement, userId)
		return true, err
	}
	return false, err
}

func (repo *UserRepositoryDB) Logout(userId int) (bool, error) {
	sqlStatement := ` UPDATE users SET logout_time = NOW() login_time = NULL WHERE id = $1;`
	_, err := repo.DB.Exec(sqlStatement, userId)
	if err == nil {
		return true, nil
	}
	return false, err
}

func (repo *UserRepositoryDB) Delete(userId int) error {
	sqlStatement := ` DELETE users WHERE id = $1;`
	_, err := repo.DB.Exec(sqlStatement, userId)
	if err == nil {
		return nil
	}
	return err
}
