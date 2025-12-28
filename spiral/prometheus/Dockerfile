FROM golang:1.25.5-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /home
COPY go.mod go.sum ./
RUN go mod download
COPY spiral/prometheus spiral/prometheus
RUN go build -o main spiral/prometheus/*.go

FROM alpine:3.17 AS final
COPY --from=builder /home/main main
EXPOSE 8000
CMD ["./main"]

