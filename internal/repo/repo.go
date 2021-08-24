package repo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-question-api/internal/models"
)

const tableName = "questions"

var ErrorNotFound = errors.New("not found")

// Repo - интерфейс хранилища для сущности Question
type Repo interface {
	AddEntity(ctx context.Context, entity *models.Question) error
	AddEntities(ctx context.Context, entities []models.Question) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Question, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Question, error)
	RemoveEntity(ctx context.Context, entityId uint64) error
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
	query := squirrel.Insert(tableName).
		Columns("user_id", "text").
		Values(entity.UserId, entity.Text).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&entity.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) AddEntities(ctx context.Context, entities []models.Question) error {
	query := squirrel.
		Insert(tableName).
		Columns("user_id", "text").
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, entity := range entities {
		query = query.Values(entity.UserId, entity.Text)
	}

	_, err := query.QueryContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Question, error) {
	query := squirrel.Select("id", "user_id", "text").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

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
	query := squirrel.Select("id", "user_id", "created", "link").
		From(tableName).
		Where(squirrel.Eq{"id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

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

func (r *repo) RemoveEntity(ctx context.Context, entityId uint64) error {
	query := squirrel.Delete(tableName).
		Where(squirrel.Eq{"id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

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
