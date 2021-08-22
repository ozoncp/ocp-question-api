package api

import (
	"context"
	desc "github.com/ozoncp/ocp-question-api/pkg/ocp-question-api"

	"github.com/rs/zerolog/log"
)

type questionApiServer struct {
	desc.UnimplementedOcpQuestionApiServer
}

func NewOcpQuestionApiServer() desc.OcpQuestionApiServer {
	return &questionApiServer{}
}

func (s *questionApiServer) CreateQuestionV1(
	context context.Context,
	request *desc.CreateQuestionV1Request,
) (*desc.CreateQuestionV1Response, error) {
	log.Print("The question was created successful")
	return &desc.CreateQuestionV1Response{}, nil
}

func (s *questionApiServer) DescribeQuestionV1(
	context context.Context,
	request *desc.DescribeQuestionV1Request,
) (*desc.DescribeQuestionV1Response, error) {
	log.Printf("Reading the question #%v was successful", request.GetQuestionId())
	return &desc.DescribeQuestionV1Response{}, nil
}

func (s *questionApiServer) ListQuestionsV1(
	context context.Context,
	request *desc.ListQuestionsV1Request,
) (*desc.ListQuestionsV1Response, error) {
	log.Print("Reading of the questions was successful")
	return &desc.ListQuestionsV1Response{}, nil
}

func (s *questionApiServer) RemoveQuestionV1(
	context context.Context,
	request *desc.RemoveQuestionV1Request,
) (*desc.RemoveQuestionV1Response, error) {
	log.Printf("Removing of the question #%v was successful", request.GetQuestionId())
	return &desc.RemoveQuestionV1Response{}, nil
}
