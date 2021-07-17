package utils

import (
	"errors"
)

func Split(source []int, batchSize int) ([][]int, error) {
	if batchSize <= 0 {
		return nil, errors.New("batch size cannot be less than or equal to zero")
	}

	batchCount := len(source)/batchSize + 1
	result := make([][]int, 0)

	var first, last int

	for i := 0; i < batchCount; i++ {
		first = i * batchSize
		last = first + batchSize

		if last > len(source) {
			last = len(source)
		}

		if first == last {
			break
		}

		result = append(result, source[first:last])
	}

	return result, nil
}
