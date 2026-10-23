package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("n должно быть больше 0")
		return
	}
	var sum int
	var avg float64
	maxValue := math.MinInt
	minValue := math.MaxInt
	var countEven int
	var countOdd int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
		maxValue = max(maxValue, a)
		minValue = min(minValue, a)
		if a%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	avg = float64(sum) / float64(n)
	fmt.Println(
		"Сумма:", sum, "\n",
		"Среднее", avg, "\n",
		"Максимальное:", maxValue, "\n",
		"Минимальное:", minValue, "\n",
		"Количество четных:", countEven, "\n",
		"Количество нечетных:", countOdd)
}
