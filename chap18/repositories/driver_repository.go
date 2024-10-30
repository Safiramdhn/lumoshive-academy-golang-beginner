package repositories

import (
	"database/sql"
	"golang-beginner-18/models"
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

func (repo *DriverRepositoryDB) GetActiveDriversByMonth() (*[]models.OrderSummary, error) {
	sqlStatement := `
	SELECT DATE_TRUNC('month', o.order_date) AS month, d.id, CONCAT(d.first_name, ' ', d.last_name) AS driver_full_name,
 		COUNT(o.driver_id) AS total_order
	FROM orders o
	JOIN driver d ON o.driver_id = d.id
	GROUP BY month, d.id, driver_full_name
	ORDER BY month, total_order DESC;`

	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var drivers []models.OrderSummary
	for rows.Next() {
		var driver models.OrderSummary
		err := rows.Scan(&driver.Month, &driver.Id, &driver.Name, &driver.TotalOrders)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return &drivers, nil
}
func (repo *DriverRepositoryDB) CountDriverLogin() (int, int, error) {
	// Get the number of active driver and the total number of driver
	sqlStatement := `SELECT COUNT(d.id) AS total_driver_login
	FROM driver d
	JOIN users u ON d.user_id = u.id
	WHERE u.login_time IS NOT NULL;`

	row := repo.DB.QueryRow(sqlStatement)
	var totalDriverLogin int
	var err error
	if err = row.Scan(&totalDriverLogin); err != nil {
		return 0, 0, err
	}
	sqlStatement = `SELECT COUNT(c.id) AS total_driver_login
	FROM driver c
	JOIN users u ON c.user_id = u.id
	WHERE u.logout_time IS NOT NULL;`

	row = repo.DB.QueryRow(sqlStatement)
	var totalDriverLogout int
	if err = row.Scan(&totalDriverLogout); err != nil {
		return 0, 0, err
	}
	return totalDriverLogin, totalDriverLogout, nil
}
