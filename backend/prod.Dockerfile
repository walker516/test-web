FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /server .

COPY ./config.dev.json /root/config.json
COPY ./templates /root/templates

EXPOSE 8080

CMD ["./server"]
