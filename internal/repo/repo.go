package repo

import "github.com/ozoncp/ocp-question-api/internal/models"

// Repo - интерфейс хранилища для сущности Question
type Repo interface {
	AddEntities(entities []models.Question) error
	ListEntities(limit, offset uint64) ([]models.Question, error)
	DescribeEntity(entityId uint64) (*models.Question, error)
}
