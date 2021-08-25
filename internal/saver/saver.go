package saver

import (
	"context"
	"log"
	"time"

	"github.com/ozoncp/ocp-question-api/internal/flusher"
	"github.com/ozoncp/ocp-question-api/internal/models"
)

type Saver interface {
	Init()
	Save(entity models.Question)
	Close()
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	duration time.Duration,
) Saver {
	entities := make(chan models.Question, capacity)
	done := make(chan struct{})

	return &saver{
		entities: entities,
		done:     done,
		flusher:  flusher,
		duration: duration,
	}
}

type saver struct {
	entities chan models.Question
	done     chan struct{}
	flusher  flusher.Flusher
	duration time.Duration
}

func (s *saver) Init() {
	go func() {
		entities := make([]models.Question, 0)
		ticker := time.NewTicker(s.duration)

		defer func() {
			ticker.Stop()
			s.flush(entities)
			close(s.entities)
		}()

		for {
			select {
			case entity := <-s.entities:
				entities = append(entities, entity)
			case <-ticker.C:
				entities = s.flush(entities)
			case <-s.done:
				return
			}
		}
	}()
}

func (s *saver) Save(entity models.Question) {
	s.entities <- entity
}

func (s *saver) Close() {
	s.done <- struct{}{}
}

func (s *saver) flush(entities []models.Question) []models.Question {
	result, err := s.flusher.Flush(context.TODO(), entities)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
