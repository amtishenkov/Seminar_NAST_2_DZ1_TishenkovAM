package main

import "fmt"

func main() {
	var a float64
	var b float64
	var ch string
	fmt.Scan(&a, &b, &ch)
	switch ch {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль!")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
