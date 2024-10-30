package repositories

import (
	"database/sql"
	"golang-beginner-18/models"
	"time"
)

type CustomerRepositoryDB struct {
	DB *sql.Tx
}

func NewCustomerRepositoryDB(db *sql.Tx) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{DB: db}
}

func (repo *CustomerRepositoryDB) Create(customer *models.Customer) error {
	// Create a new customer in the database
	sqlStatement := `INSERT INTO customer (user_id, first_name, last_name) VALUES ($1, $2, $3);`
	_, err := repo.DB.Exec(sqlStatement, customer.UserId, customer.FirstName, customer.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CustomerRepositoryDB) GetById(id int) (*models.Customer, error) {
	// Get a customer by id from the database
	sqlStatement := `SELECT id, first_name, last_name, user_id FROM customer WHERE id = $1;`
	row := repo.DB.QueryRow(sqlStatement, id)
	customer := &models.Customer{}
	err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.UserId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (repo *CustomerRepositoryDB) GetAll() (*[]models.Customer, error) {
	sqlStatement := `SELECT id, first_name, last_name, user_id FROM customer;`
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := []models.Customer{}
	for rows.Next() {
		customer := &models.Customer{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.UserId)
		if err != nil {
			return nil, err
		}
		customers = append(customers, *customer)
	}
	return &customers, nil
}

func (repo *CustomerRepositoryDB) CheckCustomer() (int, int, error) {
	// Get the number of active customers and the total number of customers
	sqlStatement := `SELECT COUNT(c.id) AS total_customer_login
	FROM customer c
	JOIN users u ON c.user_id = u.id
	WHERE u.login_time IS NOT NULL;`

	row := repo.DB.QueryRow(sqlStatement)
	var totalCustomerLogin int
	var err error
	if err = row.Scan(&totalCustomerLogin); err != nil {
		return 0, 0, err
	}
	sqlStatement = `SELECT COUNT(c.id) AS total_customer_login
	FROM customer c
	JOIN users u ON c.user_id = u.id
	WHERE u.logout_time IS NOT NULL;`

	row = repo.DB.QueryRow(sqlStatement)
	var totalCustomerLogout int
	if err = row.Scan(&totalCustomerLogout); err != nil {
		return 0, 0, err
	}
	return totalCustomerLogin, totalCustomerLogout, nil
}

func (repo *CustomerRepositoryDB) GetFrequentCustomersByMonth(startDate, endDate time.Time) ([]interface{}, error) {
	sqlStatement := `SELECT concat(c.first_name, ' ', c.last_name) AS customer_name,
	(
		SELECT
			COUNT(O.ID)
		FROM
			ORDERS O
		WHERE
			O.customer_id = c.ID 
			AND O.ORDER_TIME BETWEEN $1 AND $2
	)
	FROM customer c`

	rows, err := repo.DB.Query(sqlStatement, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []interface{}
	for rows.Next() {
		customer := struct {
			Name       string
			OrderCount int
		}{}
		err := rows.Scan(&customer.Name, &customer.OrderCount)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
