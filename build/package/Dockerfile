FROM golang:1.18.1-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod tidy && go build -o feishu main.go

FROM alpine:3.7

COPY --from=builder /build/feishu /usr/local/bin/feishu
RUN chmod +x /usr/local/bin/feishu
