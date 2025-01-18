package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func fetchWeather(city string, apiKey string, ch chan<- string, wg *sync.WaitGroup) interface{} {
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

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("TOKEN")
	cities := []string{"Toronto", "Denver", "London", "Tokyo", "Ankara"}
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		fmt.Printf("Getting weather for %s \n", city)
		go fetchWeather(city, apiKey, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

}
