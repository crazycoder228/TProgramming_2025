package main

import (
	"fmt"
	"math"
)

func main() {
	// Вывод результата выполнения функции calculate с точностью до 4 знаков после запятой
	fmt.Printf("y = %.4f\n", calculate(2.5, 4.6, 1.2))
}

func calculate(a, b, x float64) float64 {
	// Вычисление (a + b*x)
	abx := a + b*x

	// Вычисление числителя: (a + b*x)^2.5
	numerator := math.Pow(abx, 2.5)

	// Вычисление знаменателя: 1 + log10(a + b*x)
	denominator := 1 + math.Log10(abx)

	// Итоговый результат
	y := numerator / denominator

	// Возврат резултата
	return y
}