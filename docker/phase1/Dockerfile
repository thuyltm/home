FROM golang:1.25.5-alpine
WORKDIR /home
COPY go.mod go.sum ./
RUN go mod download
COPY phase1 phase1
RUN go build -o main ./phase1
CMD ["./main"]