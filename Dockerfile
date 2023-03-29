
FROM golang:1.20.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main /app/main

WORKDIR /app

# Expose port 8000
# EXPOSE 8000

CMD ["/app/main"]


