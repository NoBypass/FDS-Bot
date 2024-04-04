FROM golang:1.22.1-alpine
LABEL authors="NoBypass"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bot ./cmd/bot/main.go

ENTRYPOINT ["/app/bot"]