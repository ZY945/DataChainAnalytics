-- 创建数据库
CREATE DATABASE IF NOT EXISTS blockchain_master CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE blockchain_master;

-- 区块链信息配置表
CREATE TABLE IF NOT EXISTS chain_configs (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '链名称',
    chain_id INT NOT NULL COMMENT '链ID',
    chain_type VARCHAR(50) NOT NULL COMMENT '链类型(eth/bsc等)',
    native_symbol VARCHAR(20) NOT NULL COMMENT '原生代币符号',
    rpc_urls JSON NOT NULL COMMENT 'RPC地址列表',
    browser_urls JSON NOT NULL COMMENT '浏览器地址',
    icon_url VARCHAR(255) COMMENT '链图标URL',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
    priority INT NOT NULL DEFAULT 0 COMMENT '优先级',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_name` (`name`),
    UNIQUE KEY `uk_chain_id` (`chain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区块链信息配置表';

-- 采集任务配置表
CREATE TABLE IF NOT EXISTS collection_tasks (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    chain_id BIGINT UNSIGNED NOT NULL COMMENT '关联的链ID',
    name VARCHAR(100) NOT NULL COMMENT '任务名称',
    task_type VARCHAR(50) NOT NULL COMMENT '任务类型(event/tx/block)',
    description TEXT COMMENT '任务描述',
    contract_address VARCHAR(42) COMMENT '合约地址',
    abi_content TEXT COMMENT '合约ABI',
    start_height BIGINT UNSIGNED NOT NULL COMMENT '起始高度',
    end_height BIGINT UNSIGNED COMMENT '结束高度(0表示持续同步)',
    current_height BIGINT UNSIGNED COMMENT '当前处理高度',
    collection_rules JSON NOT NULL COMMENT '采集规则配置',
    concurrent_num INT NOT NULL DEFAULT 1 COMMENT '并发数',
    batch_size INT NOT NULL DEFAULT 100 COMMENT '批次大小',
    interval_ms INT NOT NULL DEFAULT 5000 COMMENT '采集间隔(毫秒)',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-停止 1-运行 2-暂停',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_name` (`name`),
    KEY `idx_chain_id` (`chain_id`),
    FOREIGN KEY (`chain_id`) REFERENCES chain_configs(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采集任务配置表';

-- Webhook配置表
CREATE TABLE IF NOT EXISTS webhooks (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT 'webhook名称',
    url VARCHAR(255) NOT NULL COMMENT '回调地址',
    secret_key VARCHAR(255) COMMENT '密钥',
    task_ids JSON NOT NULL COMMENT '关联的任务ID列表',
    retry_times INT NOT NULL DEFAULT 3 COMMENT '重试次数',
    retry_interval INT NOT NULL DEFAULT 1000 COMMENT '重试间隔(毫秒)',
    timeout_ms INT NOT NULL DEFAULT 5000 COMMENT '超时时间(毫秒)',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Webhook配置表';

-- 告警规则表
CREATE TABLE IF NOT EXISTS alert_rules (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '规则名称',
    rule_type VARCHAR(50) NOT NULL COMMENT '规则类型(system/business)',
    description TEXT COMMENT '规则描述',
    conditions JSON NOT NULL COMMENT '告警条件',
    notification_channels JSON NOT NULL COMMENT '通知渠道配置',
    silence_minutes INT NOT NULL DEFAULT 60 COMMENT '静默时间(分钟)',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警规则表';

-- 告警记录表
CREATE TABLE IF NOT EXISTS alert_records (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    rule_id BIGINT UNSIGNED NOT NULL COMMENT '关联的规则ID',
    alert_level VARCHAR(20) NOT NULL COMMENT '告警级别(info/warning/error/critical)',
    title VARCHAR(255) NOT NULL COMMENT '告警标题',
    content TEXT NOT NULL COMMENT '告警内容',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态: 0-未处理 1-已处理 2-已忽略',
    handle_time TIMESTAMP NULL COMMENT '处理时间',
    handle_user VARCHAR(100) COMMENT '处理人',
    handle_note TEXT COMMENT '处理说明',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `idx_rule_id` (`rule_id`),
    KEY `idx_status` (`status`),
    FOREIGN KEY (`rule_id`) REFERENCES alert_rules(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警记录表';

-- 系统日志表
CREATE TABLE IF NOT EXISTS system_logs (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    log_type VARCHAR(50) NOT NULL COMMENT '日志类型(system/business/security)',
    level VARCHAR(20) NOT NULL COMMENT '日志级别(debug/info/warn/error)',
    module VARCHAR(50) NOT NULL COMMENT '模块名称',
    action VARCHAR(100) NOT NULL COMMENT '操作名称',
    content TEXT NOT NULL COMMENT '日志内容',
    operator VARCHAR(100) COMMENT '操作人',
    ip_address VARCHAR(50) COMMENT 'IP地址',
    user_agent VARCHAR(255) COMMENT '用户代理',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `idx_log_type` (`log_type`),
    KEY `idx_level` (`level`),
    KEY `idx_module` (`module`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统日志表';

-- 插入示例数据
INSERT INTO chain_configs (name, chain_id, chain_type, native_symbol, rpc_urls, browser_urls, priority) VALUES
('Ethereum', 1, 'eth', 'ETH', 
 '["https://eth-mainnet.alchemyapi.io/v2/your-api-key", "https://mainnet.infura.io/v3/your-api-key"]',
 '["https://etherscan.io"]', 
 100),
('BSC', 56, 'bsc', 'BNB', 
 '["https://bsc-dataseed1.binance.org", "https://bsc-dataseed2.binance.org"]',
 '["https://bscscan.com"]', 
 90);

-- 插入告警规则示例
INSERT INTO alert_rules (name, rule_type, description, conditions, notification_channels) VALUES
('系统CPU告警', 'system',
 'CPU使用率超过阈值告警',
 '{
     "metric": "cpu_usage",
     "operator": ">",
     "threshold": 80,
     "duration": "5m"
 }',
 '{
     "telegram": {
         "chat_id": "your-chat-id",
         "token": "your-bot-token"
     },
     "email": ["admin@example.com"]
 }'
);

-- 插入Webhook示例
INSERT INTO webhooks (name, url, task_ids, retry_times, retry_interval) VALUES
('数据同步回调', 'https://api.example.com/webhook/sync',
 '[1, 2]',
 3, 1000
);

-- 配置表
CREATE TABLE `api_config` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(64) NOT NULL COMMENT 'API名称',
  `url` varchar(255) NOT NULL COMMENT 'API地址',
  `token` varchar(255) DEFAULT NULL COMMENT '认证token',
  `secret` varchar(255) DEFAULT NULL COMMENT '密钥',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT 'API类型: 1-WebSocket 2-HTTP 3-其他',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-停用 1-启用',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0-否 1-是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='API配置表'; 

INSERT INTO api_config (name, url, token, type) VALUES 
('alltick_gold', 'wss://quote.tradeswitcher.com/quote-b-ws-api', 'your-token', 1); 