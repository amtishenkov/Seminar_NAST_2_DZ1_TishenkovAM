package main

import "fmt"

func main() {
	var n int
	var sum int
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}
	fmt.Print(sum)
}
