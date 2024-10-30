package repositories

import (
	"golang-beginner-18/models"
	"time"
)

type UserRepository interface {
	Create(first_name, last_name, email, password string) (int, error)
	Login(email, password string) (bool, error)
	Logout(userId int) error
	Delete(userId int) error
}

type CustomersRepository interface {
	Create(customer *models.Customer) error
	GetAll() (*[]models.Customer, error)
	CheckCustomer() (int, int, error)
	GetById(id int) (*models.Customer, error)
	GetFrequentCustomersByMonth(startDate, endDate time.Time) ([]interface{}, error)
}

type DriverRepository interface {
	Create(driver *models.Driver) error
	GetAll() (*[]models.Driver, error)
	GetById(id int) (*models.Driver, error)
	GetActiveDriversByMonth(startDate, endDate time.Time) ([]interface{}, error)
}

type OrderRepository interface {
	Create(order *models.Orders) error
	GetAll(order *[]models.Orders) (*[]models.Orders, error)
	GetById(id int) (*models.Orders, error)
	GetTotalOrdersByMonth() (*[]models.Orders, error)
	GetPopularAreas() (*[]models.Orders, error)
	GetOrderPeakHours() (*[]models.Orders, error)
}
