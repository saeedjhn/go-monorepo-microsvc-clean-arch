FROM golang:1.23.2 AS builder

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR $PROJECT_DIR

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM golang:1.23.2-alpine3.20 AS production

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR $PROJECT_DIR

EXPOSE 8000 50051

CMD ["build/user/main"]
