package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-beginner-19/models"
	"strconv"
	"strings"
)

type MentorRepositoryDB struct {
	DB *sql.DB
}

func NewMentorRepositoryDB(db *sql.DB) *MentorRepositoryDB {
	return &MentorRepositoryDB{DB: db}
}

func (repo *MentorRepositoryDB) Update(mentor *models.Mentor) error {
	fields := make(map[string]interface{})

	if mentor.FirstName != "" {
		fields["first_name"] = mentor.FirstName
	}
	if mentor.LastName != "" {
		fields["last_name"] = mentor.LastName
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

	sqlStatement := fmt.Sprintf("UPDATE mentor SET %s updated_at = NOW() WHERE id = $%d", strings.Join(setClauses, ", "), index)
	values = append(values, mentor.ID)
	_, err := repo.DB.Exec(sqlStatement, values...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MentorRepositoryDB) Delete(id int) error {
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

	sqlStatement := "UPDATE mentor SET status = 'deleted' WHERE id = $1 RETURNING user_id"
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

func (repo *MentorRepositoryDB) GetById(id int) (*models.Mentor, error) {
	sqlStatement := "SELECT id, first_name, last_name, status FROM mentor WHERE id = $1"
	row := repo.DB.QueryRow(sqlStatement, id)
	mentor := &models.Mentor{}

	err := row.Scan(&mentor.ID, &mentor.FirstName, &mentor.LastName, &mentor.Status)
	if err == sql.ErrNoRows {
		return nil, errors.New("mentor not found")
	} else if err != nil {
		return nil, err
	}

	return mentor, nil
}

func (repo *MentorRepositoryDB) GetAll() (*[]models.Mentor, error) {
	sqlStatement := `SELECT id, first_name, last_name, status FROM mentor WHERE status = 'active'`
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mentors := []models.Mentor{}
	for rows.Next() {
		mentor := &models.Mentor{}
		err := rows.Scan(&mentor.ID, &mentor.FirstName, &mentor.LastName, &mentor.Status)
		if err != nil {
			return nil, err
		}
		mentors = append(mentors, *mentor)
	}
	return &mentors, nil
}
