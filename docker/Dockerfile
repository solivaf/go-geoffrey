FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
VOLUME /app/go-geoffrey
COPY . /app/go-geoffrey
WORKDIR /app/go-geoffrey
RUN go install github.com/solivaf/go-geoffrey

FROM alpine
COPY --from=builder /go/bin/go-geoffrey /go/bin/go-geoffrey