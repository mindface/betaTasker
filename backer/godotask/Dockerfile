FROM golang:1.21

RUN mkdir -p /usr/local/go/godotask
WORKDIR /usr/local/go/godotask

COPY go.mod .
COPY go.sum .

RUN go mod download

# アプリケーションコードをコピー
COPY ./ .