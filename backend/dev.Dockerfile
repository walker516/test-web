FROM golang:1.24-alpine AS dev

WORKDIR /app

RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest


COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
