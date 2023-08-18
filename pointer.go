package main

import "fmt"

func interact(num *int) {
	*num = 10000
}

func main() {
	myValue := 300
	fmt.Print(myValue)
	interact(&myValue)
	fmt.Print("\n", myValue)
}
