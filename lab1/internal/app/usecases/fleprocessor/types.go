package fleprocessor

import (
	"lab1/internal/app/domain"
)

type SimpleProcessorFunc func(domain.CalculationValues) (domain.CalculationValues, error)

type FileWriter interface {
	WriteValues(values domain.CalculationValues) error
}

type FileReader interface {
	ReadValues() (domain.CalculationValues, error)
}
