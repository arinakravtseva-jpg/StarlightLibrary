package main

import "fmt"

func main() {
	libraryName := "Starlight Library"
	bookCount := 120
	rating := 4.8
	isOpen := true

	bookTitle := "A Court of Thorns and Roses"
	isAvailable := true

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
}
