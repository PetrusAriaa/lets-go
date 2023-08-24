package main

import (
	"fmt"
	"math/rand"
	"time"
)

func find(upperBound, lowerBound, randInt, guess, n_step int) int {
	guess = lowerBound + ((upperBound - lowerBound) / 2)
	if guess == randInt {
		fmt.Println("Found!")
		return n_step
	}
	if randInt < guess {
		upperBound = guess
		n_step++
	} else if randInt > guess {
		lowerBound = guess
		n_step++
	}
	return find(upperBound, lowerBound, randInt, guess, n_step)
}

func main() {
	size := 10000
	randInt := rand.Intn(size)
	fmt.Printf("randomized number: %d\n", randInt)
	start := time.Now()
	res := find(size, 0, randInt, size+1, 0)
	fmt.Printf("Time step: %d\n", res)
	fmt.Println("Time taken:", start.Sub(time.Now()))
}
