package filewriter

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

func NewJSON(fileName string) *JSON {
	return &JSON{
		filePath: fileName,
	}
}

func (w *JSON) WriteValues(values domain.CalculationValues) error {
	file, err := os.Create(w.filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(values); err != nil {
		return fmt.Errorf("json.Encode: %w", err)
	}

	return nil
}
