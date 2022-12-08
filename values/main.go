package main

import (
	"fmt"
)

func main() {
	// strings

	fmt.Println("Melvin " + "raj")

	//integers
	fmt.Println("Sum of 10 and 20:", 10+20)
	fmt.Println("Multiplication of 10 and 20:", 10*20)
	fmt.Println("Multiplication of 10 and 20:", 10-20)
	fmt.Println("Multiplication of 10 and 20:", 10/30)
	// 0.5 is 0 in int.

	// floor is the neaarest bottom value of the given number
	// ceiling is the nearest top value of the given number.
	// 10.7 ceil(10.7) - 11.
	// 10.7 fllor(10.7) - 10.

	//float
	fmt.Println("Sum of 10 and 20:", 10.0+2.6)
	fmt.Printf("Multiplication of 10 and 20:%.2f\n", 10.0*20.0)
	fmt.Printf("Multiplication of 10 and 20: %.2f\n", 10.0-20.0)
	fmt.Printf("Multiplication of 10 and 20:%.2f\n", 10.0/30.0)

	// boolean

	// true and true is true.
	// false and false is false.
	// false or false is false.
	// true or true is true.

	fmt.Println(true && true)
	fmt.Println(false && false)
	fmt.Println(false || false)
	fmt.Println(!true)

	//variables

	var name string = "Melvin"
	fmt.Println(name)
	fmt.Printf("type of name is %T is ", name)

	var firstName, lastName string = "Melvin", "Raj"

	fmt.Println(firstName, lastName)
	fmt.Printf("type of firstName and lastName is %T, %T\n", firstName, lastName)
	fmt.Println("-------------------------")
	var decimal = 10.5
	decimal = 10.6
	// once we declare var we can't redeclare it.

	fmt.Println(decimal)
	fmt.Printf("type of decimal is %T\n", decimal)
	fruits := "apple"
	fmt.Printf("type of fruits is %T", fruits)

}
