package filereader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"lab1/internal/app/domain"
)

type JSON struct {
	filePath string
}

func NewJSON(filePath string) *JSON {
	return &JSON{
		filePath: filePath,
	}
}

func (j *JSON) ReadValues() (domain.CalculationValues, error) {
	file, err := os.Open(j.filePath)
	if err != nil {
		return domain.CalculationValues{}, fmt.Errorf("os.Open: %w", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	values := new(domain.CalculationValues)
	if err = json.NewDecoder(file).Decode(values); err != nil {
		return domain.CalculationValues{}, err
	}

	return *values, nil
}
