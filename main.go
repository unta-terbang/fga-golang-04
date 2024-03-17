package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Status struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

func DetermineStatus(water, wind int) (string, string) {
	var statusWater, statusWind string

	switch {
	case water < 5:
		statusWater = "Aman"
	case water >= 6 && water <= 8:
		statusWater = "Siaga"
	default:
		statusWater = "Bahaya"
	}

	switch {
	case wind < 6:
		statusWind = "Aman"
	case wind >= 7 && wind <= 15:
		statusWind = "Siaga"
	default:
		statusWind = "Bahaya"
	}

	return statusWater, statusWind
}

func UpdateJSON(filePath string) {
	for {
		
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		statusWater, statusWind := DetermineStatus(water, wind)

		status := Status{
			Water:       water,
			Wind:        wind,
			StatusWater: statusWater,
			StatusWind:  statusWind,
		}

		jsonData, err := json.MarshalIndent(status, "", "    ")
		if err != nil {
			panic(err)
		}

		err = WriteToFile(filePath, jsonData)
		if err != nil {
			panic(err)
		}

		fmt.Println("Data Elemen Air dan Angin:", string(jsonData))

		time.Sleep(15 * time.Second)
	}
}

func WriteToFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	filePath := "data_wawi.json"

	go UpdateJSON(filePath)

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		var status Status
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&status)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(status)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonData)
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
