package filewriter

import (
	"lab1/internal/app/domain"
)

type JSONMock struct {
	writtenValues domain.CalculationValues
}

func NewJSONMock() *JSONMock {
	return &JSONMock{}
}

func (j *JSONMock) WriteValues(values domain.CalculationValues) error {
	j.writtenValues = values

	return nil
}

func (j *JSONMock) GetLastWrittenValues() domain.CalculationValues {
	return j.writtenValues
}
