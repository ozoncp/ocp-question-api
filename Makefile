#!make
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: run
run:
	GOBIN=$(LOCAL_BIN) go run cmd/ocp-question-api/main.go

.PHONY: lint
lint:
	GOBIN=$(LOCAL_BIN) golint ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build: .vendor-proto .generate .build

.PHONY: .build
.build:
	GOBIN=$(LOCAL_BIN) go build -o $(LOCAL_BIN)/ocp-question-api cmd/ocp-question-api/main.go

.PHONY: generate
generate: .vendor-proto .generate

.PHONY: .generate
.generate:
	mkdir -p swagger
	mkdir -p pkg/ocp-question-api
	PATH="$(PATH):$(LOCAL_BIN)" protoc -I vendor.protogen \
		--go_out=pkg/ocp-question-api --go_opt=paths=import \
		--go-grpc_out=pkg/ocp-question-api --go-grpc_opt=paths=import \
		--grpc-gateway_out=pkg/ocp-question-api \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=import \
		--swagger_out=allow_merge=true,merge_file_name=api:swagger \
		--validate_out lang=go:pkg/ocp-question-api \
		api/ocp-question-api/ocp-question-api.proto
	mv pkg/ocp-question-api/github.com/ozoncp/ocp-question-api/pkg/ocp-question-api/* pkg/ocp-question-api/
	rm -rf pkg/ocp-question-api/github.com
	mkdir -p cmd/ocp-question-api
	cd pkg/ocp-question-api && ls go.mod || go mod init github.com/ozoncp/ocp-question-api/pkg/ocp-question-api && go mod tidy

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api/ocp-question-api
	cp api/ocp-question-api/ocp-question-api.proto vendor.protogen/api/ocp-question-api/ocp-question-api.proto
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
		mkdir -p vendor.protogen/github.com/envoyproxy &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
	fi

.PHONY: deps
deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
	@if [ ! -f go.mod ]; then \
		go mod init github.com/ozoncp/ocp-question-api ;\
	fi
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto@v1.5.2
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc@v1.40.0
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	GOBIN=$(LOCAL_BIN) go get -u github.com/envoyproxy/protoc-gen-validate@v0.6.1
	GOBIN=$(LOCAL_BIN) go get -u github.com/rs/zerolog@v1.23.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate


.PHONY: migrate
migrate: .install-migrate-deps .migrate

.PHONY: .install-migrate-deps
.install-migrate-deps:
	go get -u github.com/pressly/goose/v3/cmd/goose@v3.1.0

.PHONY: .migrate
.migrate:
	goose -dir database/migrations postgres "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_EXTERNAL_HOST}:${DB_EXTERNAL_PORT}/${DB_DATABASE}?sslmode=disable" up
