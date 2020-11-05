FROM golang:1.14-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver

FROM alpine:3.9 

COPY --from=builder /app/httpserver /app/
EXPOSE 8090
ENTRYPOINT ["/app/httpserver"]
