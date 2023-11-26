FROM golang:1.21

LABEL version="1.0"
LABEL name="forum"
LABEL description="Live forum"
LABEL maintainer="tvooglai"

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY forum.sql .
COPY example.sql .
COPY ui ./ui
COPY handlers ./handlers
COPY api ./api
RUN go build -o bin .

EXPOSE 8080

ENTRYPOINT [ "/app/bin" ]
