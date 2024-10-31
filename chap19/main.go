/*
buatkan design tabel Aplikasi E-learning dengan fitur app sebagai berikut :
- user terdiri dari (admin, siswa, mentor)
- admin dapat mengimput data siswa, data mentor, materi, kelas, jadwal, pengumuman
- siswa dapat melihat jadwal kelas, materi (video), absensi, nilai pribadi, tugas perminggu
- mentor dapat welihat jadwal kelas, materi (video), absensi siswanya, absensi mentor, tugas yang dibuat
siswa dan mentbr dapat melihat reader board (siswa yang terbaik) dan pengumuman

latihan peraktik hari ini :
- implementasi dari soal di atas khusus role admin
- implementasi login sebelum mengelola data
- admin bisa mengelola (CRUD) data siswa, mentor, materi, jadwal dan pengumuman
*/

package main

import (
	"fmt"
	"golang-beginner-19/configs"
	"golang-beginner-19/handlers"
	"log"
)

func main() {
	db, err := configs.InitDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	var endpoint string
	fmt.Println("endpoint: ")
	fmt.Scan(&endpoint)

	switch endpoint {
	case "admin-login":
		handlers.AdminLoginHandler(db)
	case "create-student":
		handlers.CreateStudentHandler(db)
	case "update-student":
		handlers.UpdateStudentHandler(db)
	case "delete-student":
		handlers.DeleteStudentHandler(db)
	case "get-student-by-id":
		handlers.GetStudentByIdHandler(db)
	case "get-all-students":
		handlers.GetAllStudentsHandler(db)
	case "update-user":
		handlers.UpdateUserHandler(db)
	case "create-mentor":
		handlers.CreateMentorHandler(db)
	case "get-mentor-by-id":
		handlers.GetMentorByIdHandler(db)
	case "get-all-mentors":
		handlers.GetAllMentorsHandler(db)
	case "update-mentor":
		handlers.UpdateMentorHandler(db)
	case "delete-mentor":
		handlers.DeleteMentorHandler(db)
	case "create-material":
		handlers.CreateMaterialHandler(db)
	case "get-material-by-id":
		handlers.GetMaterialByIdHandler(db)
	case "get-all-materials":
		handlers.GetAllMaterialsHandler(db)
	default:
		fmt.Println("Invalid endpoint")
	}
}
