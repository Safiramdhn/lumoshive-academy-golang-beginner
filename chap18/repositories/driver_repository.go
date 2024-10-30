package repositories

import (
	"database/sql"
	"golang-beginner-18/models"
	"time"
)

type DriverRepositoryDB struct {
	DB *sql.Tx
}

func NewDriverRepositoryDB(db *sql.Tx) *DriverRepositoryDB {
	return &DriverRepositoryDB{DB: db}
}

func (repo *DriverRepositoryDB) Create(driver *models.Driver) error {
	sqlStatement := `INSERT INTO driver (user_id, first_name, last_name) VALUES ($1, $2, $3);`
	_, err := repo.DB.Exec(sqlStatement, driver.Id, driver.FirstName, driver.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (repo *DriverRepositoryDB) GetById(id int) (*models.Driver, error) {
	// Get a customer by id from the database
	sqlStatement := `SELECT id, first_name, last_name, user_id FROM driver WHERE id = $1;`
	row := repo.DB.QueryRow(sqlStatement, id)
	driver := &models.Driver{}
	err := row.Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.UserId)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

func (repo *DriverRepositoryDB) GetAll() (*[]models.Driver, error) {
	sqlStatement := `SELECT id, first_name, last_name, user_id FROM driver;`
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	drivers := []models.Driver{}
	for rows.Next() {
		driver := &models.Driver{}
		err := rows.Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.UserId)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, *driver)
	}
	return &drivers, nil
}

func (repo *DriverRepositoryDB) GetActiveDriversByMonth(startDate, endDate time.Time) ([]interface{}, error) {
	sqlStatement := `
	SELECT D.ID, CONCAT(D.FIRST_NAME, ' ', D.LAST_NAME) AS DRIVER_FULL_NAME,
	(
		SELECT
			COUNT(O.ID)
		FROM
			ORDERS O
		WHERE
			O.DRIVER_ID = D.ID 
			AND O.ORDER_TIME BETWEEN $1 AND $2
	)
	FROM DRIVER D`
	rows, err := repo.DB.Query(sqlStatement, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var drivers []interface{}
	for rows.Next() {
		driver := struct {
			Id             int
			DriverFullName string
			TotalOrders    int
		}{}
		err := rows.Scan(&driver.Id, &driver.DriverFullName, &driver.TotalOrders)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}
