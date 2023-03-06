package fleprocessor

import (
	"lab1/internal/app/usecases/filereader"
	"lab1/internal/app/usecases/filewriter"

	"github.com/p-alexander/configurator"
)

type simpleConfig struct {
	fileReader FileReader
	fileWriter FileWriter
}

func newDefaultSimpleConfig() *simpleConfig {
	return &simpleConfig{
		fileReader: filereader.NewJSON("input.json"),
		fileWriter: filewriter.NewJSON("output.json"),
	}
}

func WithFileReader(fileReader FileReader) configurator.Option[*simpleConfig] {
	return func(c *simpleConfig) error {
		c.fileReader = fileReader

		return nil
	}
}

func WithFileWriter(fileWriter FileWriter) configurator.Option[*simpleConfig] {
	return func(c *simpleConfig) error {
		c.fileWriter = fileWriter

		return nil
	}
}
