package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNum := rand.Intn(10)
	var guess int
	fmt.Println("You have 5 chance to guess number between 0 and 10")
	for i := 0; i < 5; i++ {
		fmt.Print("Guess: ")
		fmt.Scan(&guess)
		if i == 4 && guess != randomNum {
			fmt.Println("Try again :(")
			break
		}
		if guess > randomNum {
			fmt.Println("Guess lower")
		} else if guess < randomNum {
			fmt.Println("Guess higher")
		} else if guess == randomNum {
			fmt.Printf("Gotcha %v", guess)
			break
		}

	}

}
