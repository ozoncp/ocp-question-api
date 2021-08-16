package flusher

import (
	"errors"
	"github.com/ozoncp/ocp-question-api/internal/models"
	"github.com/ozoncp/ocp-question-api/internal/repo"
	"github.com/ozoncp/ocp-question-api/internal/utils"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(entities []models.Question) ([]models.Question, error)
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize uint,
	entityRepo repo.Repo,
) (Flusher, error) {
	if chunkSize == 0 {
		return nil, errors.New("batch size cannot be equal to zero")
	}

	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}, nil
}

type flusher struct {
	chunkSize  uint
	entityRepo repo.Repo
}

func (f *flusher) Flush(entities []models.Question) ([]models.Question, error) {
	bulks, err := utils.SplitToBulks(entities, f.chunkSize)
	if err != nil {
		return entities, err
	}

	for i, bulk := range bulks {
		if err := f.entityRepo.AddEntities(bulk); err != nil {
			return entities[uint(i)*f.chunkSize:], err
		}
	}

	return nil, nil
}
