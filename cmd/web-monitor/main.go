package main

import (
	"fmt"
	"os"
	"net/http"
)

const (
	START        = "1- Start Monitoring"
	SHOW         = "2- Show Logs"
	EXIT         = "0- Exit Program"
	MISTAKE = "Inespected command..."
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
		fmt.Println(err)
	}
	fmt.Println("O comando escolhido foi", input)
}

func managementSelection() {
	switch input {
	case "1":
		Monitoring()
	case "2":
		DisplayLogs()
	case "0":
		ExitProgram()
	default:
		fmt.Println(MISTAKE)
		os.Exit(-1)
	}
}

func Monitoring() {
	fmt.Println("Monitoring...")
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-OK response status:", resp.Status)
		return
	} else {
		fmt.Println("Successfully connected to the URL.", "Status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
}

func DisplayLogs() {
	fmt.Println("Displaying logs...")
}

func ExitProgram() {
	fmt.Println("Exiting program...")
	os.Exit(0)
}
