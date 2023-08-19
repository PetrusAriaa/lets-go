package main

import "fmt"

func bubbleSortAscending(arr *[]int, counter *int) {
	var buffer int
	unsortedArray := *arr
	count := *counter
	isSorted := true
	for i := 0; i < len(unsortedArray); i++ {
		if i != len(unsortedArray)-1 && unsortedArray[i] > unsortedArray[i+1] {
			isSorted = false
			buffer = unsortedArray[i]
			unsortedArray[i] = unsortedArray[i+1]
			unsortedArray[i+1] = buffer
		}
	}
	fmt.Printf("Run[%d]: \t%v\n", count, unsortedArray)

	if !isSorted {
		count++
		bubbleSortAscending(&unsortedArray, &count)
	}
}

func bubbleSortDescending(arr *[]int, counter *int) {
	var buffer int
	unsortedArray := *arr
	count := *counter
	isSorted := true
	for i := 0; i < len(unsortedArray); i++ {
		if i != len(unsortedArray)-1 && unsortedArray[i] < unsortedArray[i+1] {
			isSorted = false
			buffer = unsortedArray[i]
			unsortedArray[i] = unsortedArray[i+1]
			unsortedArray[i+1] = buffer
		}
	}
	fmt.Printf("Run[%d]: \t%v\n", count, unsortedArray)

	if !isSorted {
		count++
		bubbleSortDescending(&unsortedArray, &count)
	}
}

func main() {
	unsortedArray := []int{3, 4, 1, 1, 9, 5, 3, 7, 4, 5, 8, 8, 1, 10, 5, 4, 3, 7, 6, 8, 9}
	counter := 1
	bubbleSortAscending(&unsortedArray, &counter)
}
