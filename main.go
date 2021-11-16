package main

import (
	"fmt"
	"module07/calc"
)

func main() {
	var firstNumber, secondNumber float64
	var operator string
	var err error

	fmt.Println("Hello, your upgraded structure based calculator is here, please enter first number and hit Enter")
	if _, err = fmt.Scanln(&firstNumber); err != nil {
		fmt.Println("Error reading", err)
	}
	fmt.Println("What is the operation ou wish to perform? Currently supported options are (+,-,*,/). Don't forget to hit Enter")
	if _, err = fmt.Scanln(&operator); err != nil {
		fmt.Println("Error reading", err)
	}
	fmt.Println("Enter second number and hit Enter")
	if _, err = fmt.Scanln(&secondNumber); err != nil {
		fmt.Println("Error reading", err)
	}

	calculation := calc.NewCalculator()

	result := calculation.Calculate(firstNumber, secondNumber, operator)

	fmt.Printf("Result is: %.2f\n", result)
}
