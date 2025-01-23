package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var calculator string

	fmt.Scan(&num1, &num2, &calculator)

	switch calculator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	case "%":
		if int(num2) == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			result := int(num1) % int(num2)
			fmt.Printf("Результат: %d\n", result)
		}
	default:
		fmt.Println("ошибка")
	}
}
