#!/bin/bash

# 创建主目录结构
mkdir -p backend/cmd/{collector,analyzer}
mkdir -p backend/internal/{collector/{service,config},analyzer/{service,config},api}
mkdir -p backend/pkg/{persistence/{mysql,redis},notification/{telegram,discord},types,utils}
mkdir -p backend/configs
mkdir -p frontend/{src,public}
mkdir -p docs

# 创建基本文件
touch backend/go.mod
touch backend/go.sum
touch frontend/package.json
touch backend/cmd/collector/main.go
touch backend/cmd/analyzer/main.go

# 创建内部包文件
touch backend/internal/collector/service/collector.go
touch backend/internal/collector/config/config.go
touch backend/internal/analyzer/service/analyzer.go
touch backend/internal/analyzer/config/config.go
touch backend/internal/api/api.go

# 创建公共包文件
touch backend/pkg/persistence/mysql/mysql.go
touch backend/pkg/persistence/redis/redis.go
touch backend/pkg/notification/telegram/telegram.go
touch backend/pkg/notification/discord/discord.go
touch backend/pkg/types/types.go
touch backend/pkg/utils/utils.go

# 创建配置文件
touch backend/configs/config.yaml 