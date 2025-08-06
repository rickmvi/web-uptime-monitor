package main

import (
	"fmt"
)

const (
	START        = "1- Start Monitoring"
	SHOW         = "2- Show Logs"
	EXIT         = "3- Exit Program"
	TO_CONTINUED = "To continued..."
)

var (
	input string
)

func main() {
	var name string
	version := "1.0.0"

	fmt.Println("Hello, sr. Please enter your name:")
	fmt.Print("Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hello, sr.", name)
	fmt.Println("The program is in version", version)

	menu()
	selectMenu()
	managementSelection()
}

func menu() {
	fmt.Println(START)
	fmt.Println(SHOW)
	fmt.Println(EXIT)
}

func selectMenu() {

	fmt.Println("Selecione um comando:")
	_, err := fmt.Scan(&input)
	if err != nil {
		err.Error()
	}
	fmt.Println("O comando escolhido foi", input)
}

func managementSelection() {
	switch input {
	case "1":
		fmt.Println(START)
	case "2":
		fmt.Println(SHOW)
	case "3":
		fmt.Println(EXIT)
	default:
		fmt.Println(TO_CONTINUED)
	}
}
