package main

import "fmt"

func main() {
	var balance int
	fmt.Scan(&balance)
	//Конец - просмотр баланса
	for {
		fmt.Println(
			"1. Посмотреть баланс - нажмите 1",
			"2. Пополнить баланс - нажмите 2",
			"3. Снять средства - нажмите 3")
		var ch int
		var sum int
		fmt.Scan(&ch)
		switch ch {
		case 1:
			fmt.Println(balance)
			fmt.Println("Всего хорошего!")
			return
		case 2:
			fmt.Println("Введите вносимую сумму")
			fmt.Scan(&sum)
			if sum <= 0 {
				fmt.Println("Сумма должна быть положительной")
			} else {
				balance += sum
			}
		case 3:
			fmt.Println("Введите сумму, которую хотите снять")
			fmt.Scan(&sum)
			if sum <= 0 {
				fmt.Println("Сумма должна быть положительной")
			} else if sum > balance {
				fmt.Println("Недостаточно средств на счете")
			} else {
				balance -= sum
			}
		default:
			fmt.Println("Некорректная опция, повторите попытку")
		}
	}

}
