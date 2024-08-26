package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Maps Module To Find Your Self Easily")

	language := make(map[string]string)

	language["en"] = "English"
	language["es"] = "Español"
	language["de"] = "Deutsch"
	language["it"] = "Italiano"
	language["ar"] = "العربية"

	fmt.Println("Languages are :", language)

	// Find the language by key
	for key, val := range language {
		fmt.Printf("for key %v Value is %v\n", key, val)
	}
	// Update a language
	language["ar"] = "العربية (Arabic)"
	fmt.Println("Updated Languages are :", language)

	// Delete a language
	delete(language, "de")
	fmt.Println("Languages after deletion are :", language)

}
