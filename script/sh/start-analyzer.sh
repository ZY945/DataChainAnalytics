#!/bin/bash

# 设置工作目录为项目根目录
cd "$(dirname "$0")/../.." || exit

echo "启动分析器服务..."

# 确保配置文件存在
if [ ! -f "backend/configs/analyzer.yaml" ]; then
    echo "错误: 配置文件 backend/configs/analyzer.yaml 不存在"
    exit 1
fi

# 确保编译目录存在
mkdir -p backend/bin

# 编译analyzer服务
cd backend || exit
go build -o bin/analyzer cmd/analyzer/main.go

if [ $? -ne 0 ]; then
    echo "编译失败"
    exit 1
fi

# 运行服务
./bin/analyzer -config configs/analyzer.yaml 