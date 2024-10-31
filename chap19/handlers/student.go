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

func CreateStudentHandler(db *sql.DB) {
	userInput := struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{}

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

		if err := decoder.Decode(&userInput); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	userService := services.UserService{RepoUser: repositories.UserRepositoryDB{DB: db}}
	err = userService.CreateUser(userInput.Email, userInput.Password, userInput.FirstName, userInput.LastName, "student")

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error creating student, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Student created successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func UpdateStudentHandler(db *sql.DB) {
	var student models.Student
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

		if err := decoder.Decode(&student); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	studentService := services.StudentService{RepoStudent: repositories.StudentRepositoryDB{DB: db}}
	err = studentService.UpdateStudent(&student)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error update student, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Student updated successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func DeleteStudentHandler(db *sql.DB) {
	var student models.Student
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

		if err := decoder.Decode(&student); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	studentService := services.StudentService{RepoStudent: repositories.StudentRepositoryDB{DB: db}}
	err = studentService.DeleteStudent(student.ID)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error deleting student, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Student deleted successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func GetStudentByIdHandler(db *sql.DB) {
	var student models.Student
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

		if err := decoder.Decode(&student); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	studentService := services.StudentService{RepoStudent: repositories.StudentRepositoryDB{DB: db}}
	studentFound, err := studentService.GetStudentById(student.ID)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error student not found, %v", err)
		response = models.Response{StatusCode: 404, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "OK", Data: studentFound}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func GetAllStudentsHandler(db *sql.DB) {
	var students []models.Student

	studentService := services.StudentService{RepoStudent: repositories.StudentRepositoryDB{DB: db}}
	students, err := studentService.GetAllStudents()

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error student not found, %v", err)
		response = models.Response{StatusCode: 404, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "OK", Data: students}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}
