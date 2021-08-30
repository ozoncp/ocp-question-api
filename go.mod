module github.com/ozoncp/ocp-question-api

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.1
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/mock v1.4.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.16.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-question-api/pkg/ocp-question-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.23.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210821163610-241b8fcbd6c8 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/ozoncp/ocp-question-api/pkg/ocp-question-api => ./pkg/ocp-question-api
