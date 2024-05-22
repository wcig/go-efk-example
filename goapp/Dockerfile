FROM golang:1.20-alpine3.17 as builder
WORKDIR /goapp
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
ADD main.go /goapp/main.go
ADD go.mod /goapp/go.mod
ADD go.sum /goapp/go.sum
RUN go build

FROM alpine:3.17
WORKDIR /goapp
RUN mkdir -p /goapp/log
COPY --from=builder /goapp/goapp .
EXPOSE 38081
EXPOSE 2112
CMD ["./goapp"]
