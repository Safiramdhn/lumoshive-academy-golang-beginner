package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-beginner-19/models"
	"golang-beginner-19/repositories"
	"golang-beginner-19/services"
	"io"
	"os"
)

func AdminLoginHandler(db *sql.DB) {
	user := models.User{}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&user); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	adminService := services.AdminService{RepoAdmin: repositories.AdminRepositoryDB{DB: db}}
	adminId, err := adminService.LoginService(user.Email, user.Password)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Account not found, %v", err)
		response = models.Response{StatusCode: 401, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Login successful", Data: adminId}
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}
