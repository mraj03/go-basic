package main

// a package is a way to group functions, and it's made up of all the files in the same directory.
// main package indicates that If a program is part of the main package,
// then go install will create a binary file; which upon execution calls main function of the program.

import "fmt"

// import is a statement where you can import different packages for others to use or to use within your code. This helps in
// structuring your code (for now.)
// ease of access to work people already have done.
// implement powerful stuff to your code.

func main() {
	// Implementing a main function to print a message to the console. A main function executes by default when you run the main package.
	// this means no matter what a go function will run only if it has a main function.
	fmt.Println("Hello world!")
	// fmt package contains functions for formatting text, including printing to the console.
	// This package is one of the standard library packages you get when you installed Go.
}
