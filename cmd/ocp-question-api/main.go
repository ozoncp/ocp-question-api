package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-question-api/internal/repo"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ozoncp/ocp-question-api/internal/api"
	desc "github.com/ozoncp/ocp-question-api/pkg/ocp-question-api"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_EXTERNAL_HOST"),
		os.Getenv("DB_EXTERNAL_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	desc.RegisterOcpQuestionApiServer(s, api.NewOcpQuestionApiServer(repo.NewRepo(db)))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpQuestionApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go runJSON()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
