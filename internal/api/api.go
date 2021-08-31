package api

import (
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/ozoncp/ocp-question-api/internal/metrics"
	"github.com/ozoncp/ocp-question-api/internal/models"
	"github.com/ozoncp/ocp-question-api/internal/producer"
	. "github.com/ozoncp/ocp-question-api/internal/repo"
	"github.com/ozoncp/ocp-question-api/internal/utils"
	desc "github.com/ozoncp/ocp-question-api/pkg/ocp-question-api"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type questionApiServer struct {
	desc.UnimplementedOcpQuestionApiServer
	repo     Repo
	producer producer.Producer
}

func NewOcpQuestionApiServer(r Repo, p producer.Producer) desc.OcpQuestionApiServer {
	return &questionApiServer{
		repo:     r,
		producer: p,
	}
}

func (s *questionApiServer) MultiCreateQuestionsV1(
	ctx context.Context,
	req *desc.MultiCreateQuestionsV1Request,
) (*desc.MultiCreateQuestionsV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Invalid arguments was received when creating a multi certificates")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateQuestionsV1")
	defer span.Finish()

	var questions []models.Question
	for _, question := range req.GetQuestions() {
		questions = append(questions, models.Question{
			UserId: question.GetUserId(),
			Text:   question.GetText(),
		})
	}

	bulks, err := utils.SplitToBulks(questions, 2)
	if err != nil {
		return nil, err
	}

	var entityIds []uint64
	for _, bulk := range bulks {
		if entities, err := s.repo.AddEntities(ctx, bulk); err != nil {
			childSpan := tracer.StartSpan("Size of data 0 bytes", opentracing.ChildOf(span.Context()))
			childSpan.Finish()

			log.Error().Err(err).Msg("Creating a questions has failed")
			return nil, err
		} else {
			childSpan := tracer.StartSpan(
				fmt.Sprintf("Size of data %d bytes", unsafe.Sizeof(bulk)),
				opentracing.ChildOf(span.Context()),
			)
			childSpan.Finish()

			for _, entity := range entities {
				entityIds = append(entityIds, entity.Id)
			}
		}
	}

	log.Info().Msgf("The questions #%v was created successful", entityIds)

	return &desc.MultiCreateQuestionsV1Response{
		QuestionIds: entityIds,
	}, nil
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
		log.Error().Err(err).Msg("Creating a question has failed")
		return nil, err
	}

	log.Info().Msgf("The question #%v was created successful", question.Id)
	metrics.CreateQuestionCounterInc()
	err = s.producer.Send(
		producer.CreateMessage(
			producer.CreateQuestion,
			producer.EventMessage{
				Id:        question.Id,
				Action:    producer.CreateQuestion.String(),
				Timestamp: time.Now().Unix(),
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("failed send message kafka")
	}

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
		log.Error().Err(err).Msgf(
			"Reading the question #%v has failed",
			req.GetQuestionId(),
		)

		return nil, err
	}

	log.Info().Msgf("Reading the question #%v was successful", req.GetQuestionId())
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

	page := req.GetPage()
	if page == 0 {
		page = 1
	}

	entities, paginator, err := s.repo.ListEntities(ctx, page)
	if err != nil {
		log.Error().Err(err).Msg("Reading of the questions has failed")
		return nil, err
	}

	log.Info().Msgf("Reading of the questions was successful")

	questions := make([]*desc.Question, 0, len(entities))

	for _, entity := range entities {
		questions = append(questions, &desc.Question{
			Id:     entity.Id,
			UserId: entity.UserId,
			Text:   entity.Text,
		})
	}

	return &desc.ListQuestionsV1Response{
		Total:       paginator.GetTotal(),
		CurrentPage: paginator.GetCurrentPage(),
		PerPage:     paginator.GetPerPage(),
		LastPage:    paginator.GetLastPage(),
		Items:       questions,
	}, nil
}

func (s *questionApiServer) UpdateQuestionV1(
	ctx context.Context,
	req *desc.UpdateQuestionV1Request,
) (*desc.UpdateQuestionV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	question := &models.Question{
		Id:     req.GetQuestionId(),
		UserId: req.GetUserId(),
		Text:   req.GetText(),
	}

	err := s.repo.UpdateEntity(ctx, question)
	if err != nil {
		log.Error().Err(err).Msgf(
			"Updating the question #%v has failed",
			req.GetQuestionId(),
		)

		return nil, err
	}

	log.Info().Msgf("Updating the question #%v was successful", req.GetQuestionId())
	metrics.UpdateQuestionCounterInc()
	err = s.producer.Send(
		producer.CreateMessage(
			producer.UpdateQuestion,
			producer.EventMessage{
				Id:        question.Id,
				Action:    producer.UpdateQuestion.String(),
				Timestamp: time.Now().Unix(),
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("failed send message kafka")
	}

	return &desc.UpdateQuestionV1Response{
		Success: true,
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
		log.Error().Err(err).Msgf(
			"Removing of the question #%v has failed",
			req.GetQuestionId(),
		)

		return nil, err
	}

	log.Info().Msgf("Removing of the question #%v was successful", req.GetQuestionId())
	metrics.RemoveQuestionCounterInc()
	err = s.producer.Send(
		producer.CreateMessage(
			producer.RemoveQuestion,
			producer.EventMessage{
				Id:        req.GetQuestionId(),
				Action:    producer.RemoveQuestion.String(),
				Timestamp: time.Now().Unix(),
			},
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("failed send message kafka")
	}

	return &desc.RemoveQuestionV1Response{
		Success: true,
	}, nil
}
