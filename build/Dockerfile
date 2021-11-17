FROM golang:1.16

WORKDIR /voz-reatime-read-chat

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod vendor

RUN CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -mod=vendor -o app cmd/main.go

#CMD ["./cli.bin", "serve"]