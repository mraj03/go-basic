package main

import "fmt"

func main() {


	for num := 1; num <= 10; num++ {
		fmt.Println(num)
	}

	fmt.Println("------------------")
	fmt.Println("loop with single statement")

	multiNum := 1

	for multiNum <= 4 {
		fmt.Println(multiNum * multiNum)
		
	}
	fmt.Println("------------------")
	fmt.Println("Condition statements - if/else")

	// checking for even numbers and printing the numbers

	i := 1

	for i = 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("Not even")
		}
	}

	fmt.Println("------------------------")
	fmt.Println("Taking a number from user and checking if the given number is a multiple of 5")

	var fiveNum int
	fmt.Print("Enter a number:\n ")
	fmt.Scanln(&fiveNum)

	if fiveNum%5 == 0 {
		fmt.Println("This number is divisible by 5")
	} else {
		fmt.Println("This number is not divisible by 5")
	}

}
