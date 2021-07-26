package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadFileInCycle(filePath string, iterations uint) error {
	readFile := func() (data []byte, err error) {
		var file *os.File

		file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
		if err != nil {
			return nil, err
		}

		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				err = closeErr
			}
		}()

		if data, err = io.ReadAll(file); err != nil {
			return nil, err
		}

		return data, err
	}

	for i := uint(0); i < iterations; i++ {
		fileData, err := readFile()

		if err != nil {
			return err
		}

		fmt.Printf(string(fileData))
	}

	return nil
}
