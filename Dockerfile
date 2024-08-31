FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o main cmd/main.go

FROM alpine AS runner

COPY --from=builder /app/main ./
COPY --from=builder /app/config.yaml ./

CMD ["./main"]