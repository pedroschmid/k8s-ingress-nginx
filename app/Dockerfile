# Build go stage
FROM golang:1.17.7-alpine AS builder
WORKDIR /application
COPY . .
RUN GO111MODULE=off go build -o main main.go

# Build alpine stage
FROM alpine:3.13
WORKDIR /application
COPY --from=builder /application/main .

# Run stage
EXPOSE 8080
CMD ["/application/main"]
