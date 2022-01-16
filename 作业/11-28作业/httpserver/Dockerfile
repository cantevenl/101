FROM golang:1.16-alpine AS build
RUN apk add --no-cache git
RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN git clone https://github.com/cantevenl/golang.git
RUN go get github.com/golang/glog
RUN go build -o /go/httpserver golang/golang-master/httpserver/main.go

FROM scratch
COPY --from=build /go/httpserver /go/httpserver
WORKDIR /go
EXPOSE 8080
ENTRYPOINT ["/go/httpserver"]