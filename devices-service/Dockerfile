FROM golang:1.20-alpine AS build

WORKDIR /app

COPY . .

RUN go run ./cmd/gen/main.go

RUN CGO_ENABLED=0 go build -o srv ./cmd/server

FROM alpine:latest

ARG PORT=8080

WORKDIR /app

COPY --from=build /app/srv /app/

EXPOSE $PORT

CMD ["/app/srv"]