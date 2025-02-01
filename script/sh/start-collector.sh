#!/bin/bash

# 检查配置文件
if [ ! -f "backend/configs/config.yaml" ]; then
    echo "配置文件不存在，从示例创建..."
    cp backend/configs/config.yaml.example backend/configs/config.yaml
    echo "请修改 backend/configs/config.yaml 中的配置"
    exit 1
fi

# 编译
cd backend
go build -o bin/collector cmd/collector/main.go

# 检查编译结果
if [ $? -ne 0 ]; then
    echo "编译失败"
    exit 1
fi

# 启动服务
echo "启动采集器服务..."
./bin/collector 