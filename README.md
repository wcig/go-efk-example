# go-efk-example

goapp -> log file -> fluent-bit -> es -> kibana

## 1. Docker compose run

```shell
docker-compose -f docker-compose.yml up
```

## 2. Kibana setting

1. Create index pattern: Setting -> Management -> Kibana -> Index patterns: Create index pattern "app_log-*".
2. View logs: Setting -> Analytics -> Discover.

## 3. Test and check logs

```shell
curl http://localhost:38081/ready
```