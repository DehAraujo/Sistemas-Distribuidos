FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
# Esta linha agora inclui todas as dependências de COMPILAÇÃO
RUN apk add --no-cache git zeromq-dev build-base g++ 
RUN go mod download
COPY . .
RUN go build -o client main.go

FROM alpine:3.18
WORKDIR /app
# ----------------------------------------------------
# CORREÇÃO AQUI: Instala a biblioteca ZeroMQ de runtime (sem -dev)
RUN apk add --no-cache zeromq 
# ----------------------------------------------------
COPY --from=build /app/client .
CMD ["./client"]