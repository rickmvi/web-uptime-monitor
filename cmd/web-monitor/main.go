package main

import (
	"fmt"
)

const (
	INIT string  = "1- Start monitoring"
	LOGS string  = "2- Exibit the logs"
	EXIT string  = "3- Exit the program"
	MSG  string  = "Hello, sr. Please enter your name:"
	ERROR string = "Invalid option, please try again."
	NAME string  = "Name: "
)

func main() {
	var name string
	var version string = "1.0.0"

	fmt.Println("Hello, sr. Please enter your name:")
	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Println("Hello, sr.", name)
	fmt.Println("The program is in version", version)

	menu()
}

func menu() {
	fmt.Println(INIT)
	fmt.Println(LOGS)
	fmt.Println(EXIT)

	var option int
	fmt.Print("Plaese, select an option: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Println("Starting monitoring...")
	case 2:
		fmt.Println("Displaying logs...")
		menu() // Call menu again to allow further actions
	case 3:
		fmt.Println("Exiting the program...")
		return // Exit the program
	default:
		fmt.Println("Invalid option, please try again.")
		menu() // Call menu again for a valid option
	}
}