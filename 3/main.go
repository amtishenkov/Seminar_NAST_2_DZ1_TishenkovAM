package main

import "fmt"

func main() {
	var n int
	var sum int
	var avg float64
	fmt.Scan(&n)
	sum = 0
	for i := 1; i < n+1; i++ {
		sum += i
	}
	avg = float64(sum) / float64(n)
	fmt.Println(sum, avg)
}
