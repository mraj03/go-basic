// let's create a CLI quiz about the tv show The office where Michael asks questions about the office as if you worked there previously.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello I am Michael Scott, Regional Manager of Dunder Mifflin. This is an interview round, You will get to work herre only if you sleep with me, No😅. Just answer the questions and you are good.")

	var answerOne string
	fmt.Println("Who is the assistant reginal manager here?")
	fmt.Scanln(&answerOne)

	if answerOne == "Dwight" || answerOne == "dwight" || answerOne == "Dwight Schrute" || answerOne == "dwight schrute" {
		fmt.Println("Okay, it is Assistant to the regional manager and you are right but wrong")
	} else {
		fmt.Println("No, you are wrong. That's what she said.")
	}

	fmt.Println("I am giving you one more chance, get it right and you will get the parking previleges.")

	var answerTwo string

	fmt.Println("Who here has agreed to do something that is very noble. That involves more than 25 students.")

	fmt.Scanln(&answerTwo)

	if answerTwo == "You" || answerTwo == "Michael Scott" || answerTwo == "michael" || answerTwo == "you" || answerTwo == "michael scott" {
		fmt.Println("My heart is very full, at the moment. You complete me.")
	} else {
		fmt.Println("Are you kidding me? Godddd! That is insulting. I did. I pledged the students that I will pay their college fee.")
	}

	fmt.Println("Anyway, thank you for coming..... That's what she said. You can start working on Monday.")

}
