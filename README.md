# go-efk-example

goapp -> log file -> fluent-bit -> es -> kibana

## 1. Docker compose 运行

```shell
docker-compose -f docker-compose.yml up
```

## 2. Kibana 设置

1. 创建 index pattern: Setting -> Management -> Kibana -> Index patterns: Create index pattern "app_log-*".
2. 查看日志: Setting -> Analytics -> Discover.

## 3. 请求接口查看日志

```shell
curl http://localhost:38081/ready
```
