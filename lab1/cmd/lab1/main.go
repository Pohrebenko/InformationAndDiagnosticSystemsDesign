package main

import (
	"fmt"
	"log"

	"lab1/internal/app/domain"
	"lab1/internal/app/usecases/calculator"
	"lab1/internal/app/usecases/filereader"
	"lab1/internal/app/usecases/filewriter"
	"lab1/internal/app/usecases/fleprocessor"
)

func main() {
	ExampleUsageScenario()
}

func ExampleUsageScenario() {
	simpleFileProcessor, err := fleprocessor.NewSimple(
		fleprocessor.WithFileReader(filereader.NewJSON("../../test/example_input.json")),
		fleprocessor.WithFileWriter(filewriter.NewJSON("../../test/example_output.json")),
	)
	if err != nil {
		panic(err)
	}

	simpleCalculator := calculator.NewSimple()

	res, err := simpleFileProcessor.Process(func(values domain.CalculationValues) (domain.CalculationValues, error) {
		sum := simpleCalculator.Add(values)
		product := simpleCalculator.Multiply(values)

		productMinusSum := simpleCalculator.Subtract(domain.CalculationValues{
			A: product,
			B: sum,
		})

		productDividedByTwo, err := simpleCalculator.Divide(domain.CalculationValues{
			A: product,
			B: 2,
		})
		if err != nil {
			return domain.CalculationValues{}, fmt.Errorf("simpleCalculator.Divide: %w", err)
		}

		return domain.CalculationValues{
			A: productMinusSum,
			B: int(productDividedByTwo),
		}, nil
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Written values are A:%d B:%d \n", res.A, res.B)
}
