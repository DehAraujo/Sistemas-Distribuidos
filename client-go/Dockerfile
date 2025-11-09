FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
# Instala libs de build (incluindo zeromq-dev e g++ para compilação)
RUN apk add --no-cache git zeromq-dev build-base g++ 
RUN go mod download
COPY . .
RUN go build -o client main.go

FROM alpine:3.18
WORKDIR /app
# ----------------------------------------------------------------------
# CORREÇÃO CRUCIAL: Instala a biblioteca ZeroMQ de runtime (SEM -dev)
RUN apk add --no-cache zeromq
# ----------------------------------------------------------------------
COPY --from=build /app/client .
CMD ["./client"]