package calculator

import (
	"errors"

	"lab1/internal/app/domain"
)

var ErrDivideByZero = errors.New("division by zero")

type Simple struct{}

func NewSimple() *Simple {
	return new(Simple)
}

func (s *Simple) Add(values domain.CalculationValues) int {
	return values.A + values.B
}

func (s *Simple) Subtract(values domain.CalculationValues) int {
	return values.A - values.B
}

func (s *Simple) Multiply(values domain.CalculationValues) int {
	return values.A * values.B
}

func (s *Simple) Divide(values domain.CalculationValues) (float64, error) {
	if values.B == 0 {
		return 0, ErrDivideByZero
	}

	return float64(values.A) / float64(values.B), nil
}
