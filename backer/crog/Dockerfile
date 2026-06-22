FROM golang:1.25-bookworm

RUN mkdir -p /usr/local/go/crog
WORKDIR /usr/local/go/crog

COPY backer/crog/go.mod .
COPY backer/crog/go.sum .
RUN go mod download || true

RUN go install github.com/air-verse/air@latest

COPY backer/crog/ .

# CMD ["go", "run", "./cmd/server"]
CMD ["air"]
