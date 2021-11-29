package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b uint64
	fmt.Fscan(os.Stdin, &a, &b)
	if a < uint64(math.Pow(2, 20)) || b < uint64(math.Pow(2, 20)) {
		fmt.Println("Ввели значения не соответствующие условию")
		return
	}
	var operation string
	_, err := fmt.Fscan(os.Stdin, &operation)
	if err != nil {
		fmt.Println("Ошибка ввода")
	}
	if operation == "*" || operation == "/" || operation == "+" || operation == "-" {
		switch operation {
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		}
	} else {
		fmt.Println("неверная операция")
	}
}
