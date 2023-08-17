package main

import (
	"fmt"
	"math"
	"math/rand"
)

const PI = math.Pi

func calcRadius(x float64, y float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2)
}

func estimateArea(validPoints, totalPoints int) float64 {
	return float64(validPoints) / float64(totalPoints)
}

func evaluatePI(area float64) {
	_pi := 4 * area
	fmt.Println("Estimated PI:", _pi)
	fmt.Println("True PI:", PI)
	fmt.Println("Squared Error:", math.Pow(_pi-PI, 2))
}

func main() {
	var x, y, r []float64
	sum := 0
	points := 1000000

	for i := 0; i < points; i++ {
		x = append(x, rand.Float64())
		y = append(y, rand.Float64())
	}

	for i := 0; i < points; i++ {
		r = append(r, calcRadius(x[i], y[i]))
	}

	for i := 0; i < points; i++ {
		if r[i] <= 1.0 {
			sum++
		}
	}

	evaluatePI(estimateArea(sum, points))

}
