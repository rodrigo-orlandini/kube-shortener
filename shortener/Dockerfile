# Builder stage
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copy go mod files and download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shortener cmd/main.go

# Runner stage
FROM alpine:3.19.0 AS runner

RUN apk add --no-cache curl

WORKDIR /app

COPY --from=builder /app/shortener .

RUN chmod +x shortener

EXPOSE 8080

CMD ["./shortener"]