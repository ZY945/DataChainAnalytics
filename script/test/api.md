# collector api
// 发送飞书消息
```
curl -X POST http://localhost:8080/api/v1/alert/feishu/send \
  -H "Content-Type: application/json" \
  -d '{
    "title": "系统通知",
    "content": "数据采集任务完成",
    "type": "info"
  }'
```

// 获取黄金价格
```
curl http://localhost:8080/api/v1/gold/price
```
