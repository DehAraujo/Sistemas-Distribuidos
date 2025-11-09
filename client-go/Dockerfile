FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN go build -o client main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/client .
CMD ["./client"]
