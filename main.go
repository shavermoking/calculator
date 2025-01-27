package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var first float64
		var second float64
		var char string

		fmt.Print("Введите первое число: ")
		_, err := fmt.Scan(&first)
		if err != nil {
			fmt.Println("Ошибка ввода первого числа")
			return
		}

		fmt.Print("Введите второе число: ")
		_, err = fmt.Scan(&second)
		if err != nil {
			fmt.Println("Ошибка ввода второго числа")
			return
		}

		fmt.Print("Введите знак (+, -, *, /, ^): ")
		_, err = fmt.Scan(&char)
		if err != nil {
			fmt.Println("Ошибка ввода знака")
			fmt.Println(err)
			return
		}

		switch char {
		case "+":
			fmt.Printf("Результат: %.2f\n", first+second)
		case "-":
			fmt.Printf("Результат: %.2f\n", first-second)
		case "*":
			fmt.Printf("Результат: %.2f\n", first*second)
		case "/":
			if second == 0 {
				fmt.Println("Ошибка: деление на ноль")
			} else {
				fmt.Printf("Результат: %.2f\n", first/second)
			}
		case "^":
			fmt.Printf("Результат: %.2f\n", math.Pow(first, second))
		default:
			fmt.Println("Неизвестная операция")
		}

		fmt.Scanln()

		fmt.Print("Хотите продолжить? (y/n): ")
		var choice string
		_, err = fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода выбора")
			return
		}

		if choice != "y" {
			fmt.Println("Завершение программы.")
			break
		}
	}
}
