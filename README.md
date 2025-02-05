# 区块链数据分析平台

## 项目结构
```
.
├── backend/ # Go 后端
│ ├── cmd/ # 主程序入口
│ │ ├── collector/ # 采集器节点启动入口
│ │ │ └── main.go
│ │ └── analyzer/ # 分析节点启动入口
│ │ └── main.go
│ ├── internal/ # 内部包
│ │ ├── collector/ # 采集器节点实现
│ │ │ ├── service/ # 采集业务逻辑
│ │ │ └── config/ # 采集器配置
│ │ ├── analyzer/ # 分析节点实现
│ │ │ ├── service/ # 分析业务逻辑
│ │ │ └── config/ # 分析器配置
│ │ └── api/ # API 接口定义
│ ├── pkg/ # 公共包
│ │ ├── persistence/ # 数据持久化模块
│ │ │ ├── mysql/ # MySQL实现
│ │ │ └── redis/ # Redis实现
│ │ ├── notification/ # 通知模块
│ │ │ ├── telegram/ # Telegram机器人
│ │ │ └── discord/ # Discord机器人
│ │ ├── types/ # 公共类型定义
│ │ └── utils/ # 工具函数
│ ├── configs/ # 配置文件目录
│ ├── go.mod
│ └── go.sum
├── frontend/ # React + Vite 前端
│ ├── src/
│ ├── public/
│ └── package.json
└── docs/ # 项目文档
```

## 启动方式

1. 启动后端
```
cd backend
go run cmd/collector/main.go
```
2. 启动前端
```
cd frontend
npm run dev
```

## 系统架构

### 后端模块

1. **采集器节点 (Collector)**
   - 职责：采集区块链上的实时数据
   - 可扩展性：支持多节点部署
   - 待定具体采集内容和策略

2. **分析节点 (Analyzer)**
   - 职责：作为主节点处理数据分析任务
   - 核心功能待定

3. **持久化模块 (Persistence)**
   - 数据库选型待定
   - 数据模型设计待定

4. **通知模块 (Notification)**
   - 支持多种机器人通知渠道
   - 告警规则配置待定

### 前端模块

- 基于 React + Vite 构建
- 具体页面和功能待定

## 开发规范

### 代码规范
- Go 代码遵循 [Effective Go](https://golang.org/doc/effective_go)
- React 代码遵循团队约定的规范（待定）

### Git 提交规范
feat: 新功能
fix: 修复
docs: 文档更改
style: 代码格式修改
refactor: 重构
test: 测试用例相关
chore: 其他修改

## 环境要求

- Go 1.20+
- Node.js 18+
- React: ^18.2.0
- React DOM: ^18.2.0
- Vite: ^4.5.0
- 其他依赖待定

## 待办事项

- [ ] 确定具体采集数据内容和范围
- [ ] 设计数据模型
- [ ] 选择合适的数据库
- [ ] 设计 API 接口
- [ ] 设计前端页面原型
- [ ] 更多待补充...

## 项目进度

项目正在初始规划阶段，后续会不断更新文档内容。