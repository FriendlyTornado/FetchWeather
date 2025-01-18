package main

import (
	"fmt"
	"os"
)

func main() {

	var apiKey = os.Environ("TOKEN")
	cities := []string{"Toronto", "Denver", "London"}

	for _, city := range cities {
		data := weather.fetchWeather(city, apiKey)
		fmt.PrintLn("Here is the weather in %s", data)
	}

}
