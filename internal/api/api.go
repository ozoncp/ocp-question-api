package api

import (
	"context"
	"github.com/ozoncp/ocp-question-api/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	. "github.com/ozoncp/ocp-question-api/internal/repo"
	desc "github.com/ozoncp/ocp-question-api/pkg/ocp-question-api"

	"github.com/rs/zerolog/log"
)

type questionApiServer struct {
	desc.UnimplementedOcpQuestionApiServer
	repo Repo
}

func NewOcpQuestionApiServer(r Repo) desc.OcpQuestionApiServer {
	return &questionApiServer{repo: r}
}

func (s *questionApiServer) CreateQuestionV1(
	ctx context.Context,
	req *desc.CreateQuestionV1Request,
) (*desc.CreateQuestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	question := &models.Question{
		UserId: req.GetUserId(),
		Text:   req.GetText(),
	}

	err := s.repo.AddEntity(ctx, question)
	if err != nil {
		log.Print("Creating a question has failed")
		return nil, err
	}

	log.Printf("The question #%v was created successful", question.Id)
	return &desc.CreateQuestionV1Response{
		QuestionId: question.Id,
	}, nil
}

func (s *questionApiServer) DescribeQuestionV1(
	ctx context.Context,
	req *desc.DescribeQuestionV1Request,
) (*desc.DescribeQuestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	entity, err := s.repo.DescribeEntity(ctx, req.GetQuestionId())
	if err != nil {
		log.Printf("Reading the question #%v has failed", req.GetQuestionId())
		return nil, err
	}

	log.Printf("Reading the question #%v was successful", req.GetQuestionId())
	return &desc.DescribeQuestionV1Response{
		Question: &desc.Question{
			Id:     entity.Id,
			UserId: entity.UserId,
			Text:   entity.Text,
		},
	}, nil
}

func (s *questionApiServer) ListQuestionsV1(
	ctx context.Context,
	req *desc.ListQuestionsV1Request,
) (*desc.ListQuestionsV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	entities, err := s.repo.ListEntities(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}

	log.Print("Reading of the questions was successful")

	questions := make([]*desc.Question, 0, len(entities))

	for _, entity := range entities {
		questions = append(questions, &desc.Question{
			Id:     entity.Id,
			UserId: entity.UserId,
			Text:   entity.Text,
		})
	}

	return &desc.ListQuestionsV1Response{
		Questions: questions,
	}, nil
}

func (s *questionApiServer) RemoveQuestionV1(
	ctx context.Context,
	req *desc.RemoveQuestionV1Request,
) (*desc.RemoveQuestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.repo.RemoveEntity(ctx, req.GetQuestionId())
	if err != nil {
		log.Printf("Removing of the question #%v has failed", req.GetQuestionId())
		return nil, err
	}

	log.Printf("Removing of the question #%v was successful", req.GetQuestionId())
	return &desc.RemoveQuestionV1Response{}, nil
}
