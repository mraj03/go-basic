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

	// program to print the numbers until the number specified by user.

	fmt.Println("Enter the number until the code should run")
	var num int
	fmt.Scanln(&num)

	for i := 1; i <=num; i++{
		fmt.Println(i)
	} 
	// program to print the numbers in reverse order until the number specified by user.

	fmt.Println("Enter the number to print in reverse")
	var num2 int
	fmt.Scanln(&num2)

	fmt.Println("The numbers in reverse are:")

	for i := num2; i >=1; i--{
		fmt.Println(i)
	} 
	// program to print the numbers until the max number and min number specified by user.

	fmt.Println("Enter the max number:")
	var maxNum int
	fmt.Scanln(&maxNum)

	fmt.Println("Enter the min number")
	var minNum int
	fmt.Scanln(&minNum)

	fmt.Println("The numbers are: ")

	for maxNum>= minNum {
		fmt.Println(maxNum)
		maxNum = maxNum -1
	} 

}
