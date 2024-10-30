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
	"golang-beginner-18/services"
	"log"
)

func main() {
	db, err := database.InitDb()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	// Create users, customers, drivers
	services.CreateUser(db, "Rian", "Klyle", "rian.klyle@example.com", "password", "customer")
}
