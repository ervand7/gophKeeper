FROM golang:1.18-alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /main server/cmd/main.go

FROM alpine:3

WORKDIR /app
COPY . .
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]