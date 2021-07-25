package utils

import (
	"errors"
	"github.com/ozoncp/ocp-question-api/internal/models"
)

func SplitToBulks(source []models.Question, batchSize uint) ([][]models.Question, error) {
	if batchSize == uint(0) {
		return nil, errors.New("batch size cannot be equal to zero")
	}

	length := uint(len(source))
	batchCount := length/batchSize + 1
	result := make([][]models.Question, 0)

	var first, last uint

	for i := uint(0); i < batchCount; i++ {
		first = i * batchSize
		last = first + batchSize

		if last > length {
			last = length
		}

		if first == last {
			break
		}

		result = append(result, source[first:last])
	}

	return result, nil
}
