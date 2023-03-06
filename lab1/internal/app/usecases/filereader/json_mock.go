package filereader

import (
	"lab1/internal/app/domain"
)

type JSONMock struct {
	values domain.CalculationValues
}

func NewJSONMock(values domain.CalculationValues) *JSONMock {
	return &JSONMock{
		values: values,
	}
}

func (j *JSONMock) ReadValues() (domain.CalculationValues, error) {
	return j.values, nil
}

func (j *JSONMock) SetA(a int) {
	j.values.A = a
}

func (j *JSONMock) SetB(b int) {
	j.values.B = b
}
