package main

import (
	"errors"
	"fmt"
	)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0.0 {
				return 0, errors.New("error: you tried to divide by zero.")
			}
	return a / b, nil
}


func main() {
	var num1 int = 4
	var num2 int = 2
	var zero int = 0

	var result int = add(num1, num2)
	fmt.Println("4 + 2 = ", result)

	result = subtract(num2, zero)
	fmt.Println("2 - 0 = ", result)

	result = multiply(num1, num2)
	fmt.Println("4 * 2 = ", result)

	result,err := divide(num1, zero)
	if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("%d / %d = %d", num1,zero,result)
			}

}
