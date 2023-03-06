package calculator_test

import (
	"errors"
	"testing"

	"lab1/internal/app/domain"
	"lab1/internal/app/usecases/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	result := calculator.NewSimple().Add(domain.CalculationValues{A: 2, B: 3})
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	result := calculator.NewSimple().Subtract(domain.CalculationValues{A: 5, B: 3})
	expected := 2

	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	result := calculator.NewSimple().Multiply(domain.CalculationValues{A: 2, B: 3})
	expected := 6

	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a, b     int
		expected float64
		err      error
	}{
		{
			name:     "valid int division",
			a:        6,
			b:        3,
			expected: 2.0,
			err:      nil,
		},
		{
			name:     "valid float division",
			a:        7,
			b:        2,
			expected: 3.5,
			err:      nil,
		},
		{
			name:     "division by zero",
			a:        6,
			b:        0,
			expected: 0,
			err:      calculator.ErrDivideByZero,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			result, err := calculator.NewSimple().Divide(domain.CalculationValues{A: tt.a, B: tt.b})
			if !errors.Is(err, tt.err) {
				t.Errorf("Divide(%d, %d) error = %v, expected error = %v", tt.a, tt.b, err, tt.err)
			}

			if result != tt.expected {
				t.Errorf("Divide(%d, %d) = %f, expected %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
