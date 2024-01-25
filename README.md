# Go-QA：Next-QA对话系统后端服务

> 前端Github仓库：[Next-QA](https://github.com/shijiahao314/next-qa)

## 技术栈

- Gin
- Gorm
- Casbin

## 部署运行

### 使用Dockerfile只构建镜像（灵活性高）

```bash
# docker build -t <image_name>:<image_tag> <path_to_dockerfile>
docker build -t go-qa:v1 .
```

### 使用Docker Compose构建整套服务

```bash
# -d 在后台启动
docker compose up -d

# 跟踪日志
docker compose logs -f <container_name>
```