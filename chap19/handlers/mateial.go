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

func CreateMaterialHandler(db *sql.DB) {
	var materialInput models.Material

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

		if err := decoder.Decode(&materialInput); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	materialService := services.MaterialService{RepoMaterial: repositories.MaterialRepositoryDB{DB: db}}
	err = materialService.CreateMaterial(&materialInput)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error creating material, %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Material created successfully", Data: nil}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func GetMaterialByIdHandler(db *sql.DB) {
	var material models.Material
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

		if err := decoder.Decode(&material); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	materialService := services.MaterialService{RepoMaterial: repositories.MaterialRepositoryDB{DB: db}}
	materialFound, err := materialService.GetMaterialById(material.ID)

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error material not found, %v", err)
		response = models.Response{StatusCode: 404, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "OK", Data: materialFound}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}

func GetAllMaterialsHandler(db *sql.DB) {
	var materials []models.Material

	materialService := services.MaterialService{RepoMaterial: repositories.MaterialRepositoryDB{DB: db}}
	materials, err := materialService.GetAllMaterials()

	var response models.Response
	if err != nil {
		errMessage := fmt.Sprintf("Error material not found, %v", err)
		response = models.Response{StatusCode: 404, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "OK", Data: materials}
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))
}
