package calculator_test

import (
	"fmt"

	"lab1/internal/app/domain"
	"lab1/internal/app/usecases/calculator"
)

func ExampleSimple_Add() {
	fmt.Println(calculator.NewSimple().Add(domain.CalculationValues{A: 2, B: 3}))
	// Output: 5
}

func ExampleSimple_Subtract() {
	fmt.Println(calculator.NewSimple().Subtract(domain.CalculationValues{A: 5, B: 3}))
	// Output: 2
}

func ExampleSimple_Multiply() {
	fmt.Println(calculator.NewSimple().Multiply(domain.CalculationValues{A: 2, B: 3}))
	// Output: 6
}

func ExampleSimple_Divide() {
	res, err := calculator.NewSimple().Divide(domain.CalculationValues{A: 6, B: 3})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	// Output: 2
}
