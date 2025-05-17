package main

import (
	"fmt"
	"math"
)

func main() {
	arraysAndFor()
}

func arraysAndFor() {
	xn, xe, dx := 1.1, 3.6, 0.5

	fmt.Println("Task A")

	for x := xn; x < xe; x += dx {
		num := (x-xn)/dx + 1

		fmt.Print("y", num)

		fmt.Println(" =", calculate(2.5, 4.6, x))
	}

	var arr [5]float64 = [5]float64{1.2, 1.28, 1.36, 1.46, 2.35}

	fmt.Println("\nTask B")

	for i := 0; i < len(arr); i++ {
		num := i + 1

		fmt.Print("y", num)

		fmt.Println(" =", calculate(2.5, 4.6, arr[i]))
	}
}

func calculate(a, b, x float64) float64 {
	abx := a + b*x

	numerator := math.Pow(abx, 2.5)

	denominator := 1 + math.Log10(abx)

	y := numerator / denominator

	return y
}
