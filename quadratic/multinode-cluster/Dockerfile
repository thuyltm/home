FROM golang:1.25.5-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /home
COPY go.mod go.sum ./
RUN go mod download
COPY quadratic/multinode-cluster quadratic/multinode-cluster
RUN go build -o main quadratic/multinode-cluster/*.go

FROM alpine:3.17 AS final
COPY --from=builder /home/main main
ENV PORT=3000
EXPOSE 3000
CMD ["./main"]