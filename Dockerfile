FROM golang:alpine AS builder

COPY . /github.com/Baraulia/AUTHORIZATION_SERVICE/
WORKDIR /github.com/Baraulia/AUTHORIZATION_SERVICE/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/Baraulia/AUTHORIZATION_SERVICE/.bin/service .
COPY --from=0 /github.com/Baraulia/AUTHORIZATION_SERVICE/configs configs/

EXPOSE 8080 8090

CMD ["./service"]
