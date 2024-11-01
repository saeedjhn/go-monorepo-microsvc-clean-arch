FROM golang:1.23.2 AS builder

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR $PROJECT_DIR

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

FROM golang:1.23.2-alpine3.20 AS development

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR $PROJECT_DIR

COPY --from=builder $PROJECT_DIR $PROJECT_DIR
COPY --from=builder /go/bin/air /usr/local/bin/air

EXPOSE 8000 50051

CMD ["air", "-c", "deployments/user/development/.air.toml"]
