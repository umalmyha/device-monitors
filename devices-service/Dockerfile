FROM golang:1.20-alpine AS build

WORKDIR /app

COPY internal .

RUN CGO_ENABLED=0 go build -o srv .

FROM alpine:latest

ARG PORT=8080

WORKDIR /app

COPY --from=build /app/srv /app/

EXPOSE $PORT

CMD ["/app/srv"]