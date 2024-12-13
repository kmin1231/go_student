FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main main.go

FROM alpine:latest
RUN apk update && apk add curl

WORKDIR /app

COPY --from=builder /app/main .

COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

ENTRYPOINT ["./main"]