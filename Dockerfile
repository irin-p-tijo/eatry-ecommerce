FROM golang:1.21.5-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .



RUN go build -v -o /app/build/api ./cmd

RUN chmod +x /app/build/api


EXPOSE 8000




CMD ["/app/build/api"]
