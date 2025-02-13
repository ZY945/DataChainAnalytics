# Collector API 使用示例

## 健康检查
```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
```

## 提交区块数据
```bash
curl -X POST http://localhost:8080/api/v1/collector/blocks \
  -H "Content-Type: application/json" \
  -d '{
    "block_number": 12345,
    "block_hash": "0x...",
    "parent_hash": "0x...",
    "timestamp": 1677649200,
    "transactions_count": 100,
    "gas_used": 1000000,
    "gas_limit": 15000000,
    "raw_data": "{\"extraData\": \"...\"}"
  }'
```

## 获取区块数据
```bash
curl http://localhost:8080/api/v1/collector/blocks/12345
```

## 提交交易数据
```bash
curl -X POST http://localhost:8080/api/v1/collector/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "tx_hash": "0x...",
    "block_number": 12345,
    "from_address": "0x...",
    "to_address": "0x...",
    "value": "1000000000000000000",
    "gas_price": 50000000000,
    "gas_used": 21000,
    "status": 1,
    "raw_data": "{\"input\": \"0x...\"}"
  }'
```

## 获取交易数据
```bash
curl http://localhost:8080/api/v1/collector/transactions/0x...
```

## 获取服务状态
```bash
curl http://localhost:8080/api/v1/collector/status
```

## 获取服务配置
```bash
curl http://localhost:8080/api/v1/collector/config
```


# collector api
// 发送飞书消息
```bash
curl -X POST http://localhost:8080/api/v1/alert/feishu/send \
  -H "Content-Type: application/json" \
  -d '{
    "title": "系统通知",
    "content": "数据采集任务完成",
    "type": "info"
  }'
```

// 获取黄金价格-text
```bash
curl http://localhost:8080/api/v1/gold/price
```

// 获取黄金价格-card
```bash
curl http://localhost:8080/api/v1/gold/alert/feishu/card
```
