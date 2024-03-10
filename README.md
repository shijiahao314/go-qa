# Go-QA：Next-QA对话系统后端服务

[![go-qa](https://goreportcard.com/badge/github.com/shijiahao314/go-qa)](https://goreportcard.com/report/github.com/shijiahao314/go-qa)

> 前端Github仓库：[Next-QA](https://github.com/shijiahao314/next-qa)

## 技术栈

- Gin
- Gorm
- Casbin

## 部署运行

### 使用 Dockerfile

```bash
# docker build -t <image_name>:<image_tag> <path_to_dockerfile>
docker run -itd --name qa-mysql --restart always -e MYSQL_ROOT_PASSWORD=qa-mysql-password -e MYSQL_DATABASE=qa -p 3306:3306 mysql
docker run -itd --name qa-redis --restart always -p 6379:6379 redis --requirepass "qa-redis-password"
docker build -t go-qa:v1 .
```

### 使用 Docker Compose

```bash
# -d 在后台启动
docker compose up -d

# 跟踪日志
docker compose logs -f <container_name>
```
