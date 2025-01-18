package main

import (
	"fmt"
	"github.com/FriendlyTornado/FetchWeather/pkg/wx"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

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
		go wx.FetchWeather(city, apiKey, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

}
