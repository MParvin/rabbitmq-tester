FROM golang:1.18.1-buster as builder

WORKDIR  /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -o rabbit-tester


FROM alpine:3.10

COPY --from=builder /app/rabbit-tester .

ENTRYPOINT ["./rabbit-tester"]
