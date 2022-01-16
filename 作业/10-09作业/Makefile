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