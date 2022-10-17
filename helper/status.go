package helper

import (
	"Assignment3-AutoReload/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

func CreateJson() {
	for {
		data := structs.Datas{}
		data.Status.Water = int(RandomNumber(0, 100))
		data.Status.Wind = int(RandomNumber(0, 100))

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal("Err: ", err.Error())
		}
		err = ioutil.WriteFile("./data.json", jsonData, 0644)

		if err != nil {
			log.Fatal("Err: ", err.Error())
		}
		time.Sleep(time.Second * 15)
	}
}

func ReloadWeb(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal("Err: ", err.Error())
	}

	var status structs.Datas

	err = json.Unmarshal(data, &status)
	if err != nil {
		log.Fatal("Err: ", err.Error())
	}

	water := status.Status.Water
	wind := status.Status.Wind
	var statusWater string
	var statusWind string

	switch {
	case water < 5:
		statusWater = "AMAN"
	case water >= 6 && status.Status.Water <= 8:
		statusWater = "SIAGA"
	case water > 8:
		statusWater = "BAHAYA"
	default:
		statusWater = "TIDAK DIKETAHUI"
	}

	switch {
	case wind < 6:
		statusWind = "AMAN"
	case wind >= 7 && status.Status.Wind <= 15:
		statusWind = "SIAGA"
	case wind > 15:
		statusWind = "BAHAYA"
	default:
		statusWind = "TIDAK DIKETAHUI"
	}

	dataStatus := map[string]interface{}{
		"waterValue":  water,
		"waterStatus": statusWater,
		"windValue":   wind,
		"windStatus":  statusWind,
	}
	fmt.Println("Water :", dataStatus["waterValue"])
	fmt.Println("Status :", dataStatus["waterStatus"])
	fmt.Println("Wind :", dataStatus["windValue"])
	fmt.Println("Status :", dataStatus["windStatus"])

	template, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal("Err :", err.Error())
	}
	template.Execute(w, dataStatus)
}
