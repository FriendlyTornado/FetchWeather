package wx

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func FetchWeather(city string, apiKey string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	type data struct {
		Weather struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching %s weather", city)
	}

	defer resp.Body.Close()

	var dt data
	if err := json.NewDecoder(resp.Body).Decode(&dt); err != nil {
		fmt.Printf("Error decoding %s weather: %s", city, err)

	}

	ch <- fmt.Sprintf("Weather is %v in %s", dt.Weather.Temp, city)

	return dt.Weather.Temp

}
