loki部署步骤

```bash
https://grafana.com/docs/loki/latest/installation/helm/

helm repo add loki https://grafana.github.io/loki/charts && helm repo update

helm pull loki/loki-stack

tar xf loki-stack-2.1.2.tgz

kubectl create ns loki

cd loki-stack/

vim values.yaml

cd /root/loki/loki-stack/charts/loki
vim values.yaml

#修改value.yaml添加数据持久化
loki.persistence.storageClass: "longhorn"
grafana.persistence.storageClass: "longhorn"

#并且添加数据生命周期，默认是永久保存，改成2周
table_manager:
retention_deletes_enabled: true
retention_period: 336h

helm upgrade --install loki -n loki .

echo $(kubectl -n loki get secret loki-grafana -o jsonpath='{.data.admin-password}') | base64 -d
jPdtzsZm3O2lH9x0gJH4juQ7ZGzwkIYPiPcpTTHB

```