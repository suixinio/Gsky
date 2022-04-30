FROM golang:1.18-alpine as builder

# GO111MODULE=on
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN set -ex \
    && apk upgrade \
    && apk add gcc libc-dev git

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
ENV GIN_MODE=release

ENTRYPOINT ["./main"]