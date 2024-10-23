/*
aplikasi pengecekan suhu
buatkan aplikasi yang memantau tiga sensor (misalnya, suhu, kelembaban, dan tekanan) menggunakan goroutine. Setiap
sensor akan mengirimkan data secara berkala melalui buffered channel, dan program harus mengambil data tersebut pada
interval tertentu menggunakan time. Ticker.

Persyaratan:
- Implementasikan 3 goroutine yang mensimulasikan pengiriman data sensor dengan buffered channel.
- Gunakan time.Ticker untuk mengambil data dori setiap sensor setiap 2 detik.
- Terapkan timeout (5 detik) menggunakan timeAfter, dan jika sensor tidak merespon dalam waktu tersebut, cetak pesan
"Sensor timeout".
- Pastikan semua sensor ditutup dengan benar dan goroutine selesai.
*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	numSensors      = 3
	tickerInterval  = 2 * time.Second
	timeoutDuration = 5 * time.Second
)

type Sensor struct {
	Name  string
	Value float64
}

func monitorSensors(wg *sync.WaitGroup, sc chan Sensor, name string, minValue, maxValue float64) {
	defer wg.Done()
	for i := 0; i < numSensors; i++ {
		data := Sensor{
			Name:  name,
			Value: minValue + rand.Float64()*(maxValue-minValue),
		}
		sc <- data
	}
	fmt.Printf("Sensor %s selesai\n", name)
}

func main() {
	sensorCh := make(chan Sensor, numSensors)
	var wg sync.WaitGroup

	runtimeCPU := runtime.NumCPU()
	fmt.Printf("CPU: %v\n", runtimeCPU)
	runtime.GOMAXPROCS(4)

	wg.Add(numSensors)
	go monitorSensors(&wg, sensorCh, "Suhu", 0, 100)
	go monitorSensors(&wg, sensorCh, "Kelembaban", 0, 100)
	go monitorSensors(&wg, sensorCh, "Tekanan", 0, 100)

	ticker := time.NewTicker(tickerInterval)
	timeout := timeoutDuration

	go func() {
		wg.Wait() // Tunggu semua goroutine selesai
		ticker.Stop()
		close(sensorCh) // Tutup channel setelah semua sensor selesai
	}()

	for {
		select {
		case <-ticker.C:
			// Menunggu data dari sensor atau timeout setelah 5 detik
			select {
			case sensorData, ok := <-sensorCh:
				if !ok {
					// Jika channel sudah ditutup, keluar dari loop
					fmt.Println("Semua sensor telah selesai diproses.")
					return
				}
				fmt.Printf("Menerima data dari sensor %s: %.2f\n", sensorData.Name, sensorData.Value)
			}
		case <-time.After(timeout):
			fmt.Println("Sensor timeout")
			return
		}
	}

	fmt.Println("Semua sensor telah selesai di proses")
}
