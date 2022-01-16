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

编译+打包+上传镜像
```bash
make build && make release && make push
```

运行
```bash
docker run -itd -p 8080:8080 --name httpserver registry.cn-chengdu.aliyuncs.com/cantevenl/httpserver:v1.0

$ curl localhost:8080/healthz
Healthz returns OK (200)
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
