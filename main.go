package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Print("Please input a command: ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "help":
			fmt.Println("This is help command.")
		case "quit":
			fmt.Println("Quit.")
			return
		default:
			fmt.Println("Wrong command!")
		}
	}
}