FROM golang:1.23-alpine3.20 AS BuildStage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN chmod +x bin/generate.sh
RUN bin/generate.sh

RUN CGO_ENABLED=0 go build -o mini-twitter

FROM alpine:latest

WORKDIR /app

COPY --from=BuildStage /app/mini-twitter /app/mini-twitter

ENTRYPOINT ["./mini-twitter"]