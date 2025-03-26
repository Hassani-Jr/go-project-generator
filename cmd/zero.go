package cmd

import (
	"fmt"
	"strconv"
)

func Add(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Print(err)
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	return fmt.Sprintf("%f + %f = %f\n", num1, num2, num1+num2)
}

func Subtract(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Print(err)
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	return fmt.Sprintf("%f - %f = %f\n", num1, num2, num1-num2)
}

func Multiply(first string, second string, shouldRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Print(err)
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2)
	}

	return fmt.Sprintf("%f", num1*num2)
}
