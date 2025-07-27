package main

import (
	"fmt"
)

const (
	VERSION float32 = 22.1
	INIT   	string  = "1- Start monitoring"
	LOGS   	string  = "2- Exibit the logs"
	EXIT   	string  = "3- Exit the program"
)

func main() {
	var name string

	fmt.Println("Hello, sr. Please enter your name:")
	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Println("Hello, sr. ", name)
	fmt.Println("The program is in version", VERSION)

	fmt.Println(INIT)
	fmt.Println(LOGS)
	fmt.Println(EXIT)

	var option int
	fmt.Print("Please, select an option: ")
	fmt.Scan(&option)

	fmt.Println()
}