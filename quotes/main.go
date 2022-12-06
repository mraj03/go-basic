package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println(quote.Hello())
	fmt.Println("----------------------")

	fmt.Println("Useful phrase for world travelers:")
	fmt.Println(quote.Glass())
	fmt.Println("----------------------")

	fmt.Println("Quote of the day:")
	fmt.Println(quote.Go())
	fmt.Println("----------------------")

	fmt.Println("An optimization truth:")
	fmt.Println(quote.Opt())
	fmt.Println("----------------------")
}
