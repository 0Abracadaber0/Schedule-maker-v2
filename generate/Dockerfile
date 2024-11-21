FROM golang:alpine AS builder

WORKDIR /builder

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o ./cmd/main ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /builder/cmd/main ./cmd/main
COPY --from=builder /builder/config/config.yaml ./config/config.yaml

EXPOSE 8088

CMD ["./cmd/main"]