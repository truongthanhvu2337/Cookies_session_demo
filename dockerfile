# Stage 1: Build app
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd

# Stage 2: Runtime 
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 3000
CMD ["./app"]