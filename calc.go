package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var calc string
	fmt.Scan(&num1, &num2, &calc)
	switch calc {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(num1 / num2)
		}
	case "%":
		intNum1 := int(num1)
		intNum2 := int(num2)
		if intNum2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(intNum1 % intNum2)
		}
	default:
		fmt.Println("ОшибОчка")
	}
}
