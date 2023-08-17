package main

import (
	"fmt"
)

func getTargetUnit() uint8 {
	var buffer uint8
	fmt.Println("Select target unit:[1-4]")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("4. Reamur")
	fmt.Print("> ")
	fmt.Scan((&buffer))

	return buffer
}

func main() {
	var currUnit uint8
	var currValue, evaluate float64
	isValid := true

	fmt.Println("Select current unit:[1-4]")
	fmt.Println("1. Celcius")
	fmt.Println("2. Fahrenheit")
	fmt.Println("3. Kelvin")
	fmt.Println("4. Reamur")
	fmt.Print("> ")
	fmt.Scan((&currUnit))

	switch currUnit {
	case 1:
		target := getTargetUnit()
		fmt.Print("Insert current value: ")
		fmt.Scan(&currValue)
		switch target {
		case 1:
			evaluate = currValue
			break
		case 2:
			evaluate = (currValue * 9 / 5) + 32
			break
		case 3:
			evaluate = currValue + 273
			break
		case 4:
			evaluate = currValue * 4 / 5
			break
		default:
			fmt.Println("Invalid input")
			isValid = false
			break
		}
		break
	case 2:
		target := getTargetUnit()
		fmt.Print("Insert current value: ")
		fmt.Scan(&currValue)
		switch target {
		case 1:
			evaluate = (currValue - 32) * 5 / 9
			break
		case 2:
			evaluate = currValue
			break
		case 3:
			evaluate = ((currValue - 32) * 5 / 9) + 273
			break
		case 4:
			evaluate = ((currValue - 32) * 5 / 9) * 4 / 5
			break
		default:
			fmt.Println("Invalid input")
			isValid = false
			break
		}
		break
	case 3:
		target := getTargetUnit()
		fmt.Print("Insert current value: ")
		fmt.Scan(&currValue)
		switch target {
		case 1:
			evaluate = currValue - 273
			break
		case 2:
			evaluate = ((currValue - 273) * 9 / 5) + 32
			break
		case 3:
			evaluate = currValue
			break
		case 4:
			evaluate = (currValue - 273) * 4 / 5
			break
		default:
			fmt.Println("Invalid input")
			isValid = false
			break
		}
		break
	case 4:
		target := getTargetUnit()
		fmt.Print("Insert current value: ")
		fmt.Scan(&currValue)
		switch target {
		case 1:
			evaluate = currValue * 5 / 4
			break
		case 2:
			evaluate = ((currValue * 5 / 4) * 9 / 5) + 32
			break
		case 3:
			evaluate = (currValue * 5 / 4) + 273
			break
		case 4:
			evaluate = currValue
			break
		default:
			fmt.Println("Invalid input")
			isValid = false
			break
		}
		break
	default:
		fmt.Println("Invalid input")
		isValid = false
	}

	if isValid {
		fmt.Printf("Target value: %.2f\n", evaluate)
	}

}
