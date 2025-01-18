package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func FetchWeather(city string, apiKey string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json: "Temp"`
		} `json: "Main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid={%s}&q=%s", apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching %s weather", city)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Error decoding %s weather", city)
	}

	return data

}
