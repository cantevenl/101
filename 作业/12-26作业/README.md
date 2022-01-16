安装istio
下载地址: https://github.com/istio/istio/releases/tag/1.12.1

```bash
#添加istioctl到PATH路径
vi ~/.bash_profile
source ~/.bash_profile

#添加istio的自动补全脚本
cp tools/istioctl.bash ~/.istioctl.bash
source ~/.istioctl.bash

#查看istio的版本号
istioctl version --remote=false

#安装install -- 确保网络插件 cni 正常安装
istioctl install --set profile=demo -y

#安装工具
cd ./istio-1.12.1/samples/addons
kubectl apply -f .
#查看安装情况
[root@master1 addons]# kubectl get po -n  istio-system
NAME                                    READY   STATUS    RESTARTS   AGE
istio-egressgateway-659cc7697b-7x9j8    1/1     Running   0          16d
istio-ingressgateway-569f64cdf8-xpzbb   1/1     Running   0          16d
istiod-85c958cd6-dv946                  1/1     Running   0          16d
jaeger-7f78b6fb65-2t5jq                 1/1     Running   0          7d4h
kiali-85c8cdd5b5-h95wb                  1/1     Running   0          14d
prometheus-69f7f4d689-ssr8d             2/2     Running   1          14d
zipkin-7fcd647cf9-zv5j6                 1/1     Running   0          7d4h

#因为没有 LoadBalancer 可以使用metalLB
kubectl create ns metallb-system
kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"
# metalb-config.yaml设置分配给LB的IP段
kubectl apply -f ./metaLB

#再次查看svc
[root@master1 metalLB]# kubectl get svc -n istio-system
NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      172.18.55.23     <none>        80/TCP,443/TCP                                                               16d
istio-ingressgateway   LoadBalancer   172.18.78.136    10.10.4.110   15021:30678/TCP,80:30184/TCP,443:30990/TCP,31400:31176/TCP,15443:31446/TCP   16d
istiod                 ClusterIP      172.18.15.183    <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        16d

#EXTERNAL-IP已经有了分配的地址，因此可以将host解析到这个地址上去对istio-egressgateway进行访问


- 卸载istio
istioctl manifest generate --set profile=demo | kubectl delete -f -
```

如何实现安全保证
七层路由规则
考虑 open tracing 的接入

将项目通过istio ingress gateway的方式发布
```bash
#istio 默认自动注入 sidecar，需要为微服务的命名空间打上标签 istio-injection=enabled
kubectl label namespace httpserver istio-injection=enabled

[root@master1 ~]# kubectl scale --replicas=0 deployment cantevenl-httpserver -n httpserver
deployment.apps/cantevenl-httpserver scaled
[root@master1 ~]# kubectl scale --replicas=1 deployment cantevenl-httpserver -n httpserver
deployment.apps/cantevenl-httpserver scaled
[root@master1 ~]# kubectl get po -n !$
kubectl get po -n httpserver
NAME                                    READY   STATUS     RESTARTS   AGE
cantevenl-httpserver-74498cf8b5-zkctp   2/2     Running   0          2m55s

#如何实现安全保证
#配置tls,并且创建路由
kubectl create -n istio-system secret tls cantevenl-tls --key=a.key --cert=a.crt
kubectl apply -f ./httpserver-gw.yaml

#七层路由规则
#创建vs
kubectl apply -f ./httpserver-vs.yaml

#考虑 open tracing 的接入
#1.golang中向下发送header时需要转换小写，否则首字母大写导致tracing失效
#2.tracing需要依赖sidecar
# 需要安装istio的tracing插件 jaeger
```
