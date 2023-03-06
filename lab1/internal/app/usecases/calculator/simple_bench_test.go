package calculator_test

import (
	"testing"

	"lab1/internal/app/domain"
	"lab1/internal/app/usecases/calculator"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.NewSimple().Add(domain.CalculationValues{A: 2, B: 3})
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.NewSimple().Subtract(domain.CalculationValues{A: 5, B: 3})
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.NewSimple().Multiply(domain.CalculationValues{A: 2, B: 3})
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = calculator.NewSimple().Divide(domain.CalculationValues{A: 6, B: 3})
	}
}
