FROM golang:1.16-alpine AS build
RUN apk add --no-cache git \
    go env -w GO111MODULE=auto \
    go env -w GOPROXY=https://goproxy.cn,direct \
    git clone https://github.com/cantevenl/golang.git \
    go get github.com/golang/glog \
    go build -o /go/httpserver golang/golang-master/httpserver/main.go

FROM scratch
COPY --from=build /go/httpserver /go/
WORKDIR /go
EXPOSE 8080
ENTRYPOINT ["/go/httpserver"]