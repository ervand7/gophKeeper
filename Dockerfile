FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./pkg ./pkg
COPY ./proto ./proto
COPY ./server ./server
RUN go build -o /main server/cmd/main.go

FROM alpine:3

WORKDIR /app
COPY ./migrations ./migrations
COPY ./cert ./cert
COPY --from=builder /main .
ENTRYPOINT ["./main"]
