package repositories

import (
	"database/sql"
	"fmt"
	"golang-beginner-18/models"
)

type OrdersRepositoryDB struct {
	DB *sql.Tx
}

func NewOrdersRepositoryDB(db *sql.Tx) *OrdersRepositoryDB {
	return &OrdersRepositoryDB{DB: db}
}

func (r *OrdersRepositoryDB) Create(order *models.Orders) (int, error) {
	sqlStatement := `
        INSERT INTO orders (customer_id, driver_id, order_date, order_time, city, district, neighborhood, street_name)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;
    `
	err := r.DB.QueryRow(
		sqlStatement,
		order.CustomerId,
		order.DriverId,
		order.OrderDate,
		order.OrderTime,
		order.City,
		order.District,
		order.Neighborhood,
		order.StreetName,
	).Scan(&order.Id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert order: %v", err)
	}
	return order.Id, nil
}

func (r *OrdersRepositoryDB) GetTotalOrder() (*[]models.OrderSummary, error) {
	sqlStatement := `
		SELECT DATE_TRUNC('month', order_date) AS month, COUNT(id) AS total_orders
		FROM orders
		GROUP BY month
		ORDER BY month DESC;
	`
	rows, err := r.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []models.OrderSummary
	for rows.Next() {
		var order models.OrderSummary
		err := rows.Scan(&order.Month, &order.TotalOrders)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return &orders, nil
}

func (r *OrdersRepositoryDB) GetPopularAreas() (*[]models.Orders, error) {
	sqlStatement := `
		SELECT city, district, neighborhood, street_name, COUNT(id) AS total_orders
		FROM orders
		GROUP BY city, district, neighborhood, street_name
		ORDER BY total_orders DESC;
	`
	rows, err := r.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []models.Orders
	for rows.Next() {
		var order models.Orders
		err := rows.Scan(&order.City, &order.District, &order.Neighborhood, &order.StreetName, &order.TotalOrders)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return &orders, nil
}

func (r *OrdersRepositoryDB) GetOrderPeakHours() ([]struct {
	HourRange   string
	TotalOrders int
}, error) {
	sqlStatement := `
		SELECT 
		TO_CHAR(order_time, 'HH24:00') || ' - ' || TO_CHAR(order_time + INTERVAL '1 hour', 'HH24:00') AS hour_range,
		COUNT(id) AS total_orders
		FROM orders
		GROUP BY hour_range
		ORDER BY total_orders DESC;
	`
	rows, err := r.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orderPeakHours []struct {
		HourRange   string
		TotalOrders int
	}
	for rows.Next() {
		fmt.Println("tes")
		resultList := struct {
			HourRange   string
			TotalOrders int
		}{}
		err := rows.Scan(&resultList.HourRange, &resultList.TotalOrders)
		if err != nil {
			return nil, err
		}
		orderPeakHours = append(orderPeakHours, resultList)
	}

	return orderPeakHours, nil
}
