package fleprocessor_test

import (
	"testing"

	"lab1/internal/app/domain"
	"lab1/internal/app/usecases/filereader"
	"lab1/internal/app/usecases/filewriter"
	"lab1/internal/app/usecases/fleprocessor"
)

func TestProcessSuccessful(t *testing.T) {
	t.Parallel()

	expectedInput := domain.CalculationValues{
		A: 10,
		B: 20,
	}
	expectedOutput := domain.CalculationValues{
		A: 15,
		B: 4,
	}

	t.Run("correct values are passed to process func", func(t *testing.T) {
		simpleFileProcessor, err := fleprocessor.NewSimple(
			fleprocessor.WithFileReader(filereader.NewJSONMock(expectedInput)),
			fleprocessor.WithFileWriter(filewriter.NewJSONMock()),
		)
		if err != nil {
			t.Fatal(err)
		}

		_, err = simpleFileProcessor.Process(func(values domain.CalculationValues) (domain.CalculationValues, error) {
			if expectedInput != values {
				t.Fatalf("Expected input: %+v, got: %+v \n", expectedInput, values)
			}

			return domain.CalculationValues{}, nil
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("correct values are passed to the filewriter", func(t *testing.T) {
		testFilewriter := filewriter.NewJSONMock()
		simpleFileProcessor, err := fleprocessor.NewSimple(
			fleprocessor.WithFileReader(filereader.NewJSONMock(expectedInput)),
			fleprocessor.WithFileWriter(testFilewriter),
		)
		if err != nil {
			t.Fatal(err)
		}

		_, err = simpleFileProcessor.Process(func(values domain.CalculationValues) (domain.CalculationValues, error) {
			return domain.CalculationValues{
				A: expectedOutput.A,
				B: expectedOutput.B,
			}, nil
		})
		if err != nil {
			t.Fatal(err)
		}

		if vals := testFilewriter.GetLastWrittenValues(); vals != expectedOutput {
			t.Fatalf("Expected output: %+v, got: %+v \n", expectedOutput, vals)
		}
	})
}
