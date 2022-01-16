编写Dockerfile文件
```bash
FROM scratch
COPY bin/httpserver /go/httpserver
WORKDIR /go
EXPOSE 8080
ENTRYPOINT ["/go/httpserver"]
```
编写Makefile
```bash
export tag=v1.0

build:
	echo "building httpserver binary"
	mkdir -p bin/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/ ./httpserver

release: build
	echo "building httpserver镜像"
	docker build -t registry.cn-chengdu.aliyuncs.com/cantevenl/httpserver:${tag} .

push: release
	echo "push镜像"
	docker push registry.cn-chengdu.aliyuncs.com/cantevenl/httpserver:${tag}
```

编译+打包
```bash
make build && make release
```

运行
```bash
docker run -itd --name httpserver -p 8080:8080 registry.cn-chengdu.aliyuncs.com/cantevenl/httpserver:v1.0
```


也可以在镜像里面多段构建
```bash
FROM golang:1.16-alpine AS build
RUN apk add --no-cache git
RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN git clone https://github.com/cantevenl/101.git
RUN cd 作业/10-09作业/
RUN make build

FROM scratch
COPY --from=build /go/httpserver /go/httpserver
WORKDIR /go
EXPOSE 8080
ENTRYPOINT ["/go/httpserver"]
```
