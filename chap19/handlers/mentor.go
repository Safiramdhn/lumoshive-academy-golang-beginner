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

func CreateMentorHandler(db *sql.DB) {
	userInput := struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		AddedBy   int    `json:"added_by"`
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
	err = userService.CreateUser(userInput.Email, userInput.Password, userInput.FirstName, userInput.LastName, "mentor", userInput.AddedBy)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error creating mentor, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Mentor created successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func UpdateMentorHandler(db *sql.DB) {
	var mentor models.Mentor
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

		if err := decoder.Decode(&mentor); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	mentorService := services.MentorService{RepoMentor: repositories.MentorRepositoryDB{DB: db}}
	err = mentorService.UpdateMentor(&mentor)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error update mentor, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Mentor updated successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func DeleteMentorHandler(db *sql.DB) {
	var mentor models.Mentor
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

		if err := decoder.Decode(&mentor); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	mentorService := services.MentorService{RepoMentor: repositories.MentorRepositoryDB{DB: db}}
	err = mentorService.DeleteMentor(mentor.ID)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error deleting mentor, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Mentor deleted successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func GetMentorByIdHandler(db *sql.DB) {
	var mentor models.Mentor
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

		if err := decoder.Decode(&mentor); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	mentorService := services.MentorService{RepoMentor: repositories.MentorRepositoryDB{DB: db}}
	mentorFound, err := mentorService.GetMentorById(mentor.ID)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error mentor not found, %v", err)
		response = models.Response{StatusCode: 404, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "OK", Data: mentorFound}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func GetAllMentorsHandler(db *sql.DB) {
	var mentors []models.Mentor

	mentorService := services.MentorService{RepoMentor: repositories.MentorRepositoryDB{DB: db}}
	mentors, err := mentorService.GetAllMentors()

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error mentor not found, %v", err)
		response = models.Response{StatusCode: 404, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "OK", Data: mentors}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}
