package main

import (
	"assignment-03/helpers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type StatusData struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var t, err = template.ParseFiles("index.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var data, waterStatus, windStatus = helpers.GetStatus()
		waterValue := strconv.Itoa(data[0])
		windValue := strconv.Itoa(data[1])
		payload := map[string]string{
			"water":  waterValue,
			"wind":   windValue,
			"waterStatus": waterStatus,
			"windStatus": windStatus,
		}
		t.Execute(w, payload)

		rawData := StatusData{}
		rawData.Status.Water = data[0]
		rawData.Status.Wind = data[1]

		jsonData, err := json.Marshal(rawData)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = os.WriteFile("status.json", jsonData, 0644)

		if err != nil {
			log.Fatal("error writing data to json file")
		}
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
