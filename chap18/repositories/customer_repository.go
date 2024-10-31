package repositories

import (
	"database/sql"
	"golang-beginner-18/models"
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

// func (repo *CustomerRepositoryDB) GetById(id int) (*models.Customer, error) {
// 	// Get a customer by id from the database
// 	sqlStatement := `SELECT id, first_name, last_name, user_id FROM customer WHERE id = $1;`
// 	row := repo.DB.QueryRow(sqlStatement, id)
// 	customer := &models.Customer{}
// 	err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return customer, nil
// }

// func (repo *CustomerRepositoryDB) GetAllUsers() (*[]models.Customer, error) {
// 	sqlStatement := `SELECT id, first_name, last_name, user_id FROM customer;`
// 	rows, err := repo.DB.Query(sqlStatement)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	customers := []models.Customer{}
// 	for rows.Next() {
// 		customer := &models.Customer{}
// 		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.UserId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		customers = append(customers, *customer)
// 	}
// 	return &customers, nil
// }

func (repo *CustomerRepositoryDB) CountCustomerLogin() (int, int, error) {
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

func (repo *CustomerRepositoryDB) GetFrequentCustomersByMonth() (*[]models.OrderSummary, error) {
	// sqlStatement := `
	// SELECT DATE_TRUNC('month', order_date) AS month, c.id, concat(c.first_name, ' ', c.last_name) AS customer_name, COUNT(o.id) AS order_count
	// FROM orders o
	// JOIN customer c ON o.customer_id = c.id
	// GROUP BY month, customer_name, c.id
	// ORDER BY month, order_count DESC;
	// `

	sqlStatement := `
	SELECT c.id,
       CONCAT(c.first_name, ' ', c.last_name) AS customer_full_name,
  ( SELECT COUNT(O.ID)
   FROM orders o
   WHERE  o.customer_id = c.id
     AND o.order_date BETWEEN '2024-10-01 00:00:00' AND '2024-10-31 23:59:59' )
FROM customer c;`

	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customerOrderSumarries []models.OrderSummary
	for rows.Next() {
		var orderSumarry models.OrderSummary

		// err := rows.Scan(&orderSumarry.Month, &orderSumarry.Id, &orderSumarry.Name, &orderSumarry.TotalOrders)
		err := rows.Scan(&orderSumarry.Id, &orderSumarry.Name, &orderSumarry.TotalOrders)
		if err != nil {
			return nil, err
		}
		customerOrderSumarries = append(customerOrderSumarries, orderSumarry)
	}
	return &customerOrderSumarries, nil
}
