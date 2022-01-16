环境
```bash
k8s集群 1.20.6
[root@master1 tekton]# kubectl get node
NAME      STATUS   ROLES                  AGE    VERSION
master1   Ready    control-plane,master   121d   v1.20.6
master2   Ready    control-plane,master   121d   v1.20.6
master3   Ready    control-plane,master   121d   v1.20.6
worker2   Ready    <none>                 121d   v1.20.6
worker3   Ready    <none>                 121d   v1.20.6
worker4   Ready    <none>                 121d   v1.20.6
worker5   Ready    <none>                 121d   v1.20.6

ceph集群
[root@master1 tekton]# ceph -s
  cluster:
    id:     6edcfa88-840b-4a99-9980-def640c4df20
    health: HEALTH_WARN
            3 daemons have recently crashed
 
  services:
    mon: 3 daemons, quorum a,b,c (age 3d)
    mgr: a(active, since 6d)
    osd: 3 osds: 3 up (since 7d), 3 in (since 5w)
    rgw: 2 daemons active (2 hosts, 1 zones)
 
  data:
    pools:   9 pools, 120 pgs
    objects: 77.97k objects, 293 GiB
    usage:   878 GiB used, 622 GiB / 1.5 TiB avail
    pgs:     120 active+clean
 
  io:
    client:   271 KiB/s wr, 0 op/s rd, 1 op/s wr

ingress-nginx
[root@master1 tekton]# kubectl get po -n ingress-nginx
NAME                             READY   STATUS    RESTARTS   AGE
ingress-nginx-controller-487lp   1/1     Running   0          121d
ingress-nginx-controller-hh5f8   1/1     Running   0          121d

```

部署
```bash
make all

[root@master1 docker-build]# kubectl get po,svc,ingress,pvc -n httpserver
NAME                                        READY   STATUS    RESTARTS   AGE
pod/cantevenl-httpserver-74498cf8b5-w6drc   1/1     Running   0          8m34s

NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/cantevenl-httpserver-svc   ClusterIP   172.18.50.240   <none>        80/TCP    14m

NAME                                                     CLASS    HOSTS           ADDRESS         PORTS     AGE
ingress.networking.k8s.io/cantevenl-httpserver-ingress   <none>   cantevenl.com   172.18.189.62   80, 443   6m14s

NAME                                   STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS      AGE
persistentvolumeclaim/log-httpserver   Bound    pvc-8e65c7de-8a1e-4e46-b7a0-0b12abf62d9d   100Mi      RWO            rook-ceph-block   7m18s
```

测试
```bash
curl --noproxy "*" -H "Host:cantevenl.com" https://$(kubectl get svc -n ingress-nginx |grep -w "ingress-nginx-controller" |awk 'NR==1{print $3}') -v -k
kubectl logs -f   cantevenl-httpserver-74498cf8b5-w6drc -n httpserver
```