package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Print("Положительное, ")
	} else if n < 0 {
		fmt.Print("Отрицательное, ")
	} else {
		fmt.Print("Ноль, ")
	}
	if n%2 == 0 {
		fmt.Print("Четное, ")
	} else {
		fmt.Print("Нечетное, ")
	}
	if n%3 == 0 && n%5 == 0 {
		fmt.Print("Делится на 3 и на 5")
	} else {
		fmt.Print("Не делится на 3 и на 5")
	}
}
