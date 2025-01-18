package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func fetchWeather(city string, apiKey string) interface{} {
	type data struct {
		Weather struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching %s weather", city)
	}

	//body, err := io.ReadAll(resp.Body)
	//json.Unmarshal(body, &data{})
	defer resp.Body.Close()

	var dt data
	if err := json.NewDecoder(resp.Body).Decode(&dt); err != nil {
		fmt.Printf("Error decoding %s weather: %s", city, err)

	}

	return dt.Weather.Temp

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("TOKEN")
	cities := []string{"Toronto", "Denver", "London"}

	for _, city := range cities {
		fmt.Printf("Getting weather for %s \n", city)
		data := fetchWeather(city, apiKey)
		fmt.Printf("Here is the weather in %v \n", data)
	}

}
