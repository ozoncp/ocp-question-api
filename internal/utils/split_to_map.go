package utils

import (
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-question-api/internal/models"
)

func SplitToMap(entities []models.Question) (map[uint64]models.Question, error) {
	output := make(map[uint64]models.Question, len(entities))

	for _, entity := range entities {
		if _, found := output[entity.Id]; found {
			return nil, errors.New(fmt.Sprintf("Duplicate Id %d", entity.Id))
		}

		output[entity.Id] = entity
	}

	return output, nil
}
