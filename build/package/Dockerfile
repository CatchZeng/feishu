FROM golang:1.16.4-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN ls -la
RUN go build -o feishu cmd/main.go

FROM alpine:3.7

COPY --from=builder /build/feishu /usr/local/bin/feishu
RUN chmod +x /usr/local/bin/feishu
