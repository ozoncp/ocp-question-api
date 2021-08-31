FROM golang:1.16-alpine AS builder
RUN apk add --update make git protoc protobuf protobuf-dev
COPY . /home/github.com/ozoncp/ocp-question-api
WORKDIR /home/github.com/ozoncp/ocp-question-api
RUN make deps && make build

FROM alpine:latest AS server
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/github.com/ozoncp/ocp-question-api/bin/ocp-question-api .
COPY --from=builder /home/github.com/ozoncp/ocp-question-api/.env .
RUN chown root:root ocp-question-api
EXPOSE 8081
EXPOSE 8082
EXPOSE 9100
CMD ["./ocp-question-api"]
