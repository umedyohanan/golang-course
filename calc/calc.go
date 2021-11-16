package calc

import (
	"fmt"
	"math"
)


//calculator is a structure to keep provided input values of 2 operands and math operation
type calculator struct {}

//NewCalculator is public method to create an instance of calculator structure
func NewCalculator() calculator {
	return calculator{}
}

//Calculate is the method to manage calculations, get 2 operands and math operation
func (c* calculator) Calculate(firstNumber float64, secondNumber float64, mathOperation string) float64 {
	switch mathOperation {
	case "+":
		return c.addition(firstNumber, secondNumber)
	case "-":
		return c.subtraction(firstNumber, secondNumber)
	case "*":
		return c.multiplication(firstNumber, secondNumber)
	case "/":
		return c.division(firstNumber, secondNumber)
	default:
		fmt.Println("Operation is not supported or is invalid, please try again")
		return math.NaN()
	}
}

//addition method, gets in input operands to calculate
func (c *calculator) addition(firstNumber float64, secondNumber float64) float64 {
	return firstNumber + secondNumber
}

//subtraction method, gets in input operands to calculate
func (c *calculator) subtraction(firstNumber float64, secondNumber float64) float64 {
	return firstNumber - secondNumber
}

//multiplication method, gets in input operands to calculate
func (c *calculator) multiplication(firstNumber float64, secondNumber float64) float64 {
	return firstNumber * secondNumber
}

//division method, gets in input operands to calculate, if second operand is zero error message is printed
func (c *calculator) division(firstNumber float64, secondNumber float64) float64 {
	if secondNumber != 0 {
		return firstNumber / secondNumber
	} else {
		fmt.Println("Division to zero is not possible, please try again with other second operand")
	}

	return math.NaN()
}