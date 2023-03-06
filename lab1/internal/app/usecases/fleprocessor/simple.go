package fleprocessor

import (
	"fmt"

	"lab1/internal/app/domain"

	"github.com/p-alexander/configurator"
)

type Simple struct {
	config *simpleConfig
}

func NewSimple(opts ...configurator.Option[*simpleConfig]) (*Simple, error) {
	c := newDefaultSimpleConfig()

	if err := configurator.Constructor[*simpleConfig](c, opts); err != nil {
		return nil, fmt.Errorf("ConfigConstructor: %w", err)
	}

	return &Simple{
		config: c,
	}, nil
}

func (s *Simple) Process(simpleProcessorFunc SimpleProcessorFunc) (domain.CalculationValues, error) {
	values, err := s.config.fileReader.ReadValues()
	if err != nil {
		return domain.CalculationValues{}, fmt.Errorf("fileReader.ReadValues: %w", err)
	}

	resValues, err := simpleProcessorFunc(values)
	if err != nil {
		return domain.CalculationValues{}, fmt.Errorf("simpleProcessorFunc: %w", err)
	}

	if err = s.config.fileWriter.WriteValues(resValues); err != nil {
		return domain.CalculationValues{}, fmt.Errorf("fileWriter.WriteValues: %w", err)
	}

	return resValues, nil
}
