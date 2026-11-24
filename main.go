package main

import "fmt"

func main() {
	libraryName := "Starlight Library"
	rating := 4.8
	isOpen := true

	books := []string{
		"A Court of Thorns and Roses",
		"A Court of Mist and Fury",
		"A Court of Wings and Ruin",
		"A Court of Frost and Starlight",
		"A Court of Silver Flames",
		"Throne of Glass",
		"Crown of Midnight",
		"The Assassin's Blade",
		"Heir of Fire",
		"Queen of Shadows",
		"Empire of Storms",
		"Tower of Dawn",
		"Kingdom of Ash",
		"House of Earth and Blood",
		"House of Sky and Breath",
		"House of Flame and Shadow",
	}
	bookCount := len(books)
	bookTitle := "A Court of Thorns and Roses"
	isAvailable := true
	bookavailability := map[string]bool{
		"A Court of Thorns and Roses": true,
		"Throne of Glass":             false,
		"House of Earth and Blood":    true,
	}
	bookavailability["Throne of Glass"] = true
	available, exists := bookavailability["Throne of Glass"]

	if exists {
		fmt.Println("Books exists in map:", available)
	} else {
		fmt.Println("Books not found in map:")
	}
	fmt.Println("Throne of Glass available:", bookavailability["Throne of Glass"])
	action := "add"

	fmt.Println("Library:", libraryName)
	fmt.Println("Books:", bookCount)
	fmt.Println("Rating:", rating)
	fmt.Println("Is open:", isOpen)

	fmt.Println("Book:", bookTitle)
	if isAvailable {
		fmt.Println("Status: Available")
	} else {
		fmt.Println("Status: Borrowed")
	}
	switch action {
	case "show":
		showBooks(books)
	case "search":
		searchBooks(books, "Throne of Glass")
	case "add":
		books = addBooks(books, "The Stolen Heir")
		showBooks(books)
	default:
		fmt.Println("Unknown action")
	}

}
func showBooks(books []string) {
	fmt.Println("Book catalog:")
	for _, book := range books {
		fmt.Println("-", book)
	}
}
func searchBooks(books []string, title string) {
	for _, book := range books {
		if book == title {
			fmt.Println("Book found:", book)
			return
		}
	}
	fmt.Println("Book not found")
}
func addBooks(books []string, title string) []string {
	books = append(books, title)
	fmt.Println("Book added:", title)
	return books
}
