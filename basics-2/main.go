package main

import (
	"fmt"
)

func main() {
	//taking input from user in CLI and printing it.
	var name string

	fmt.Printf("Enter your name: ")

	fmt.Scanln(&name)

	fmt.Printf("Hello, %v. Welcome to the Valorant Tournament for xtsi gaming!", name)
}
