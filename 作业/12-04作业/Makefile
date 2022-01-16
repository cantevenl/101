export tag=v2.0

all: build release push deployment result

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

#在k8s上部署httpserver服务
deployment:release
	kubectl create ns httpserver
	kubectl create secret tls cantevenl-tls --cert=./a.cert --key=./a.key
	kubectl apply -f ./k8s

#测试最终结果
result:
	sleep 3
	curl --noproxy "*" -H "Host:taxi-dev.eysaas.com" https://$(shell kubectl get svc -n ingress-nginx |grep -w "ingress-nginx-controller" |awk 'NR==1{print $3}') -v -k




