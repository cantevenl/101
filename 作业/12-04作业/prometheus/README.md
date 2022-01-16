prometheus安装

```bash

git clone -b release-0.8 https://github.com/prometheus-operator/kube-prometheus.git

cd /root/kube-prometheus/manifests

#Prometheus数据持久化
#添加监控数据过期时间
vim prometheus-prometheus.yaml 
spec:
  retention: 30d
  alerting:
    alertmanagers:
    - apiVersion: v2
      name: alertmanager-main
      namespace: monitoring
      port: web
  storage:
    volumeClaimTemplate:
      spec:
        storageClassName: rook-ceph-block
        accessModes: ["ReadWriteOnce"]
        resources:
          requests:
            storage: 50Gi
            
kubectl apply -f setup/
kubectl create -f setup/

kubectl get po -n monitoring

#grafana配置持久化
#首先创建grafana的pvc
vim grafana-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: grafana-pvc
  namespace: monitoring
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 20Gi
  storageClassName: csi-cephfs	#如果单节点可以使用rook-ceph-block
  
  
#修改
vim grafana-deployment.yaml
	  #- emptyDir: {}
      #  name: grafana-storage
      - name: grafana-storage
        persistentVolumeClaim:
          claimName: grafana-pvc
      - name: grafana-datasources
        secret:
          secretName: grafana-datasources
      - configMap:
          name: grafana-dashboards
        name: grafana-dashboards
      - configMap:
          name: grafana-dashboard-apiserver
        name: grafana-dashboard-apiserver
      - configMap:
          name: grafana-dashboard-controller-manager
    ........
    
 #修改镜像
 vim kube-state-metrics-deployment.yaml
image: bitnami/kube-state-metrics:2.0.0

cd /root/kube-prometheus/manifests
kubectl apply -f .

```
