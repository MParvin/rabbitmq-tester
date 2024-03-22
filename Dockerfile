FROM golang:1.18.1-buster as builder

WORKDIR  /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rabbitmq-tester .

FROM alpine:3.10

WORKDIR /app

COPY --from=builder /app/rabbit-tester .

ENTRYPOINT ["/app/rabbit-tester"]
