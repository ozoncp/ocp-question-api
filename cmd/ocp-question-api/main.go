package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozoncp/ocp-question-api/internal/api"
	"github.com/ozoncp/ocp-question-api/internal/config"
	"github.com/ozoncp/ocp-question-api/internal/db"
	"github.com/ozoncp/ocp-question-api/internal/metrics"
	"github.com/ozoncp/ocp-question-api/internal/producer"
	"github.com/ozoncp/ocp-question-api/internal/repo"
	"github.com/ozoncp/ocp-question-api/internal/tracer"
	desc "github.com/ozoncp/ocp-question-api/pkg/ocp-question-api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
)

var closers []func()

func init() {
	// load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("No .env file found")
	}
}

func run() {
	conf := config.NewConfig()

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.ExternalHost,
		conf.Database.ExternalPort,
		conf.Database.Database,
	)

	dbConn := db.Connect(dsn)
	if err != nil {
		log.Error().Err(err).Msg("failed to db connect")
	}

	s := grpc.NewServer()
	desc.RegisterOcpQuestionApiServer(s, api.NewOcpQuestionApiServer(
		repo.NewRepo(dbConn),
		producer.NewProducer(),
	))

	closers = append(closers, s.Stop)

	if err := s.Serve(listen); err != nil {
		log.Error().Err(err).Msgf("failed to serve: %v", err)
	}
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	closers = append(closers, cancel)

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

// metricsServer - metrics server
func runMetrics() {
	metrics.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		panic(err)
	}
}

func runTracer() {
	closer := tracer.InitTracer("ocp-question-api")
	closers = append(closers, func() {
		err := closer.Close()
		if err != nil {
			log.Error().Err(err).Msg("Tracer closing error")
		}
	})
}

func awaitSignal() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal...")
	<-done
	fmt.Println("exiting")

	for _, closer := range closers {
		closer()
	}
}

func main() {
	go runMetrics()
	go runJSON()
	go runTracer()
	go run()

	awaitSignal()

	os.Exit(1)
}
