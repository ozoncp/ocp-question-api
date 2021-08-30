package repo

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-question-api/internal/models"
)

const tableName = "questions"

var ErrorNotFound = errors.New("not found")

// Repo - интерфейс хранилища для сущности Question
type Repo interface {
	AddEntity(ctx context.Context, entity *models.Question) error
	AddEntities(ctx context.Context, entities []models.Question) ([]models.Question, error)
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Question, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Question, error)
	RemoveEntity(ctx context.Context, entityId uint64) error
	UpdateEntity(ctx context.Context, entity *models.Question) error
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}

type repo struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r *repo) AddEntity(ctx context.Context, entity *models.Question) error {
	query := sq.Insert(tableName).
		Columns("user_id", "text").
		Values(entity.UserId, entity.Text).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&entity.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) AddEntities(ctx context.Context, entities []models.Question) ([]models.Question, error) {
	query := sq.
		Insert(tableName).
		Columns("user_id", "text").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, entity := range entities {
		query = query.Values(entity.UserId, entity.Text)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	i := 0
	for rows.Next() {
		err = rows.Scan(&entities[i].Id)
		if err != nil {
			return nil, err
		}

		i++
	}

	return entities, nil
}

func (r *repo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Question, error) {
	query := sq.Select("id", "user_id", "text").
		From(tableName).
		Where(sq.Eq{"deleted_at": nil}).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var questions []models.Question

	for rows.Next() {
		var question models.Question

		if err := rows.Scan(
			&question.Id,
			&question.UserId,
			&question.Text,
		); err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func (r *repo) DescribeEntity(ctx context.Context, entityId uint64) (*models.Question, error) {
	query := sq.Select("id", "user_id", "text").
		From(tableName).
		Where(sq.Eq{"id": entityId}).
		Where(sq.Eq{"deleted_at": nil}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var question models.Question

	if err := query.QueryRowContext(ctx).Scan(
		&question.Id,
		&question.UserId,
		&question.Text,
	); err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *repo) UpdateEntity(ctx context.Context, entity *models.Question) error {
	query := sq.Update(tableName).
		Set("user_id", entity.UserId).
		Set("text", entity.Text).
		Where(sq.Eq{"id": entity.Id}).
		Where(sq.Eq{"deleted_at": nil}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrorNotFound
	}

	return nil
}

func (r *repo) RemoveEntity(ctx context.Context, entityId uint64) error {
	query := sq.Update(tableName).
		Set("deleted_at", time.Now()).
		Where(sq.Eq{"id": entityId}).
		Where(sq.Eq{"deleted_at": nil}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrorNotFound
	}

	return nil
}
