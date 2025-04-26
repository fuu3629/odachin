# Build Stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Run Stage
FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env  

EXPOSE 50051

ENTRYPOINT ["./main"]
