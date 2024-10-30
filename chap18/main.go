/*buatlah design struktur tabel dan query data japlikasi ojeck online dengan kebutuhan sebagai berikut :
- dapat melihat total order setiap bulan
- dapat melihat customer yang sering order tiap bulan (tampilkan namanya)
- dapat melihat daerah mana saja yang banyak ordernya
- dapat melihat pukul berapa saja order yang ramai dan sepi
- dapat melihat jumlah customer yang masih login dan logout
- dapat melihat driver yang rajin mengambil order setiap bulan

buatlah program untuk implementasi query diatas dengan menerapakan repository patter
lalu tambahkan fitur input customer , driver dan order
*/

package main

import (
	"golang-beginner-18/database"
	"golang-beginner-18/models"
	"golang-beginner-18/services"
	"log"
	"time"
)

func main() {
	db, err := database.InitDb()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	// Create users, customers, drivers
	// services.CreateUser(db, "Rian", "Klyle", "rian.klyle@example.com", "password", "customer")
	// services.CreateUser(db, "Selly", "Anna", "selly.anna@example.com", "password", "driver")

	// create orders
	// newOrder := &models.Orders{
	// 	CustomerId:   5,
	// 	DriverId:     1,
	// 	City:         "Jakarta",
	// 	District:     "Jakarta Pusat",
	// 	Neighborhood: "Cempaka Putih",
	// 	StreetName:   "Cempaka Putih",
	// 	OrderDate:    time.Now(),
	// 	OrderTime:    time.Now(),
	// }
	// services.CreateOrder(db, newOrder)

	//dapat melihat jumlah customer yang masih login dan logout
	services.CountCustomerLogin(db)

	//dapat melihat customer yang sering order tiap bulan (tampilkan namanya)
	services.GetFrequentCustomersByMonth(db)

	//--dapat melihat jumlah driver yang masih login dan logout
	services.CountDriverLogin(db)

	//dapat melihat driver yang rajin mengambil order setiap bulan
	services.GetFrequentDriversByMonth(db)

	// dapat melihat total order setiap bulan
	services.GetTotalOrder(db)

	//dapat melihat pukul berapa saja order yang ramai dan sepi
	services.GetOrderPeakHours(db)

	// dapat melihat daerah mana saja yang banyak ordernya
	services.GetPopularAreas(db)
}
