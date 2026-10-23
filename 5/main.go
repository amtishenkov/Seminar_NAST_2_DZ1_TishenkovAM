package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	maxValue := math.MinInt
	minValue := math.MaxInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		maxValue = max(maxValue, a)
		minValue = min(minValue, a)
	}
	fmt.Print(maxValue, " ", minValue)
}
