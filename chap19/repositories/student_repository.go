package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-beginner-19/models"
	"strconv"
	"strings"
)

type StudentRepositoryDB struct {
	DB *sql.DB
}

func NewStudentRepositoryDB(db *sql.DB) *StudentRepositoryDB {
	return &StudentRepositoryDB{DB: db}
}

func (repo *StudentRepositoryDB) Update(student *models.Student) error {
	fields := make(map[string]interface{})

	if student.FirstName != "" {
		fields["first_name"] = student.FirstName
	}
	if student.LastName != "" {
		fields["last_name"] = student.LastName
	}

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

	sqlStatement := fmt.Sprintf("UPDATE student SET %s updated_at = NOW() WHERE id = $%d", strings.Join(setClauses, ", "), index)
	values = append(values, student.ID)
	_, err := repo.DB.Exec(sqlStatement, values...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *StudentRepositoryDB) Delete(id int) error {
	var userId int

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

	sqlStatement := "UPDATE student SET status = 'deleted' WHERE id = $1 RETURNING user_id"
	err = tx.QueryRow(sqlStatement, id).Scan(&userId)
	if err != nil {
		return err
	}
	sqlStatement = "UPDATE users SET status = 'deleted' WHERE id = $1"
	_, err = tx.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (repo *StudentRepositoryDB) GetById(id int) (*models.Student, error) {
	sqlStatement := "SELECT id, first_name, last_name, status FROM student WHERE id = $1"
	row := repo.DB.QueryRow(sqlStatement, id)
	student := &models.Student{}

	err := row.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Status)
	if err == sql.ErrNoRows {
		return nil, errors.New("student not found")
	} else if err != nil {
		return nil, err
	}

	return student, nil
}

func (repo *StudentRepositoryDB) GetAll() (*[]models.Student, error) {
	sqlStatement := `SELECT id, first_name, last_name, status FROM student WHERE status = 'active'`
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		student := &models.Student{}
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Status)
		if err != nil {
			return nil, err
		}
		students = append(students, *student)
	}
	return &students, nil
}
