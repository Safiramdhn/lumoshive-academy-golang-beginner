package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-beginner-19/models"
	"strconv"
	"strings"
	"time"
)

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{DB: db}
}

func (repo *UserRepositoryDB) Create(email, password, first_name, last_name, role string, added_by int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
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

	var userId int
	sqlStatement := `
        INSERT INTO users (email, password)
        VALUES ($1, $2) RETURNING id
    `
	err = tx.QueryRow(sqlStatement, email, password).Scan(&userId)
	if err != nil {
		return err
	}

	switch role {
	case "admin":
		sqlStatement = `INSERT INTO admins (first_name, last_name, user_id) VALUES ($1, $2, $3);`
		_, err = tx.Exec(sqlStatement, first_name, last_name, userId)
		if err != nil {
			return err
		}
	case "student":
		sqlStatement = `INSERT INTO student (first_name, last_name, user_id, added_by) VALUES ($1, $2, $3, $4);`
		_, err = tx.Exec(sqlStatement, first_name, last_name, userId, added_by)
		if err != nil {
			return err
		}
	case "mentor":
		sqlStatement = `INSERT INTO mentor (first_name, last_name, user_id, added_by) VALUES ($1, $2, $3, $4);`
		_, err = tx.Exec(sqlStatement, first_name, last_name, userId, added_by)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid role")
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepositoryDB) Update(user *models.User) error {
	fields := make(map[string]interface{})

	if user.Email != "" {
		fields["email"] = user.Email
	}
	if user.Password != "" {
		fields["password"] = user.Password
	}
	fields["updated_at"] = time.Now()

	setClauses := []string{}
	values := []interface{}{}
	index := 1
	for field, value := range fields {
		setClauses = append(setClauses, field+"=$"+strconv.Itoa(index))
		values = append(values, value)
		index++
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	sqlStatement := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(setClauses, ", "), index)
	values = append(values, user.ID)
	_, err := repo.DB.Exec(sqlStatement, values...)
	if err != nil {
		return err
	}
	return nil
}
