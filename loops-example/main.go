package main

import (
	"fmt"
)

func main() {
	// program to showcase multiplication table of 1 to n
	var firstNum, secondNum int

	fmt.Println("Multiplication table from 1 to 3 - ")

	for firstNum = 1; firstNum <= 3; firstNum++ {
		fmt.Printf("The multiplication table of %v:\n", firstNum)
		for secondNum = 1; secondNum <= 10; secondNum++ {
			fmt.Println(firstNum, "*", secondNum, " = ", firstNum*secondNum)
		}
		fmt.Println("-----------------------------------")
	}

	// program to print the numbers in reverse

}
