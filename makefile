# Makefile for a Go project

# 设置 Go 环境变量
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODEP=$(GOCMD) get
GOINSTALL=$(GOCMD) install

# 项目名称
PROJECT_NAME=go-qa.out

# 输出的可执行文件名
EXECUTABLE=$(PROJECT_NAME)

# 构建目标
build:
	$(GOBUILD) -o $(EXECUTABLE) -v .

# 运行目标
run:
	./$(EXECUTABLE)

# 测试目标
test:
	$(GOCLEAN) -cache
	$(GOTEST) -v ./...

# 清理目标
clean:
	$(GOCLEAN)
	rm -f $(EXECUTABLE)

# 格式化代码
fmt:
	$(GOCMD) fmt ./...

# 获取依赖
get:
	$(GODEP) ./...

# 安装项目（通常用于安装可执行文件或库）
install:
	$(GOINSTALL)

# 默认目标
.PHONY: all
all: build

# 确保在并行构建时，先执行依赖获取
.PHONE: get