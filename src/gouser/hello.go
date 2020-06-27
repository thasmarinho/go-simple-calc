package main

import (
	"errors"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
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
	reader := bufio.NewReader(os.Stdin)
  fmt.Println("Simple Shell")
  fmt.Println("---------------------")

	for {
		fmt.Print("num1-> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
		num1, _ := strconv.Atoi(text)

		fmt.Print("num2-> ")
    text, _ = reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
		num2, _ := strconv.Atoi(text)

		fmt.Print("operator-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "+":
			var result int = add(num1, num2)
			fmt.Printf("%d + %d = %d\n", num1, num2, result)
		case "-":
			var result int = subtract(num1, num2)
			fmt.Printf("%d - %d = %d\n", num1, num2, result)
		case "*":
			var result int = multiply(num1, num2)
			fmt.Printf("%d * %d = %d\n", num1, num2, result)
		case "/":
			result, err := divide(num1, num2)
			if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("%d / %d = %d\n", num1,num2,result)
					}
		default:
			fmt.Println("I don't know which operation u are talking about")
		}

	}
}
