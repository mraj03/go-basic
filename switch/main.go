package main

import (
	"fmt"
)

func main(){
	// let's define the months by the number they type. 

	var month int
	fmt.Println("Enter the number of the month you were born in: ")
	fmt.Scanln(&month)

	switch month{
	case 1:
		fmt.Println("It is January")
	
	case 2:
		fmt.Println("It is February")
	
	case 3:
		fmt.Println("It is March")

	case 4:
		fmt.Println("It is April")

	case 5:
		fmt.Println("It is May")

	case 6:
		fmt.Println("It is June")

	case 7:
		fmt.Println("It is July")

	case 8:
		fmt.Println("It is August")

	case 9:
		fmt.Println("It is September")

	case 10:
		fmt.Println("It is October")

	case 11:
		fmt.Println("It is November")

	case 12:
		fmt.Println("It is December")
	// if they give wrong number, we say please provide any number of month from 1-12.
	default:
		fmt.Println("please provide any number of month from 1-12")
	}

	
	
}