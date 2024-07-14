# dev
FROM golang:1.21-alpine as dev

WORKDIR /opt/server/

RUN apk add --no-cache tzdata
# Proxy for go modules blocks
ENV GOPROXY="https://goproxy.io"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app internal/main.go

ENTRYPOINT ["/app"]


# prod