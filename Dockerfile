FROM golang:1.18-alpine AS builder
WORKDIR /app
ADD . /app
RUN go mod tidy
RUN go build -o gogin main.go

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/gogin .
EXPOSE 8080
ENTRYPOINT ./gogin
