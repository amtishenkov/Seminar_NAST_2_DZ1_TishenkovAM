package main

import "fmt"

func main() {
	var n int
	countZero := 0
	countPositive := 0
	countNegative := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a > 0 {
			countPositive++
		} else if a < 0 {
			countNegative++
		} else {
			countZero++
		}
	}
	fmt.Print(
		countPositive,
		countNegative,
		countZero)
}
