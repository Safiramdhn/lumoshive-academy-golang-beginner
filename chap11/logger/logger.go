package logger

import (
	"fmt"
	"golang-beginner-11/model"
)

func LogData(data model.Data) {
	fmt.Printf("Data ID: %d, Name: %s berhasil disimpan\n", data.Id, data.Name)
}
