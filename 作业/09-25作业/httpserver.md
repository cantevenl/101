go build -o httpserver main.go
./httpserver

#访问
$ curl localhost:8080

#返回Header、状态码、客户端ip
$ ./httpserver
os version: v0.0.1
Header key: User-Agent, Header value: curl/7.79.1
Header key: Accept, Header value: */*
Response code: 200
clientip:  127.0.0.1

$ curl localhost:8080/healthz
Healthz returns OK (200)
