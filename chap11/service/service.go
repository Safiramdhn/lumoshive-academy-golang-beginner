package service

import (
	"golang-beginner-11/logger"
	"golang-beginner-11/model"
)

func ProcessData(dataChannel chan model.Data, dataSlice *[]model.Data) {
	for data := range dataChannel {
		// Menambahkan data ke slice
		*dataSlice = append(*dataSlice, data)

		// Log data yang berhasil disimpan
		logger.LogData(data)
	}
}
