-- 创建采集器数据库
CREATE DATABASE IF NOT EXISTS collector_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE collector_db;

-- 创建区块数据表
CREATE TABLE IF NOT EXISTS block_data (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
    block_hash VARCHAR(66) NOT NULL COMMENT '区块哈希',
    parent_hash VARCHAR(66) NOT NULL COMMENT '父区块哈希',
    timestamp BIGINT UNSIGNED NOT NULL COMMENT '区块时间戳',
    transactions_count INT UNSIGNED NOT NULL COMMENT '交易数量',
    gas_used BIGINT UNSIGNED NOT NULL COMMENT '使用的gas',
    gas_limit BIGINT UNSIGNED NOT NULL COMMENT 'gas限制',
    raw_data JSON COMMENT '原始区块数据',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `idx_block_number` (`block_number`),
    UNIQUE KEY `idx_block_hash` (`block_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区块数据表';

-- 创建交易数据表
CREATE TABLE IF NOT EXISTS transaction_data (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tx_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
    block_number BIGINT UNSIGNED NOT NULL COMMENT '区块高度',
    from_address VARCHAR(42) NOT NULL COMMENT '发送地址',
    to_address VARCHAR(42) COMMENT '接收地址',
    value VARCHAR(78) NOT NULL COMMENT '交易金额',
    gas_price BIGINT UNSIGNED NOT NULL COMMENT 'gas价格',
    gas_used BIGINT UNSIGNED NOT NULL COMMENT '使用的gas',
    status TINYINT NOT NULL COMMENT '交易状态 1:成功 0:失败',
    raw_data JSON COMMENT '原始交易数据',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `idx_tx_hash` (`tx_hash`),
    KEY `idx_block_number` (`block_number`),
    KEY `idx_from_address` (`from_address`),
    KEY `idx_to_address` (`to_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='交易数据表';

-- 配置表
CREATE TABLE IF NOT EXISTS collector_configs (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL COMMENT '配置名称',
    chain_type VARCHAR(50) NOT NULL COMMENT '链类型(eth/bsc等)',
    rpc_url TEXT NOT NULL COMMENT '区块链RPC地址',
    api_key VARCHAR(255) COMMENT 'API密钥(如果需要)',
    retry_times INT NOT NULL DEFAULT 3 COMMENT '重试次数',
    retry_interval INT NOT NULL DEFAULT 1000 COMMENT '重试间隔(毫秒)',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-禁用 1-启用',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采集器配置表';

-- 采集任务表
CREATE TABLE IF NOT EXISTS collector_tasks (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    config_id BIGINT UNSIGNED NOT NULL COMMENT '关联的配置ID',
    name VARCHAR(100) NOT NULL COMMENT '任务名称',
    description TEXT COMMENT '任务描述',
    contract_address VARCHAR(42) COMMENT '合约地址',
    start_block BIGINT UNSIGNED COMMENT '起始区块',
    end_block BIGINT UNSIGNED COMMENT '结束区块(0表示持续同步)',
    collection_fields JSON NOT NULL COMMENT '采集字段配置',
    task_interval INT NOT NULL DEFAULT 5000 COMMENT '任务间隔(毫秒)',
    batch_size INT NOT NULL DEFAULT 100 COMMENT '批次大小',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '状态: 0-停止 1-运行 2-暂停',
    last_block BIGINT UNSIGNED COMMENT '最后处理的区块',
    error_count INT NOT NULL DEFAULT 0 COMMENT '错误计数',
    last_error TEXT COMMENT '最后一次错误信息',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_name` (`name`),
    KEY `idx_config_id` (`config_id`),
    FOREIGN KEY (`config_id`) REFERENCES collector_configs(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采集任务表';

-- 采集数据表
CREATE TABLE IF NOT EXISTS collected_data (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    task_id BIGINT UNSIGNED NOT NULL COMMENT '关联的任务ID',
    block_number BIGINT UNSIGNED NOT NULL COMMENT '区块号',
    block_hash VARCHAR(66) NOT NULL COMMENT '区块哈希',
    transaction_hash VARCHAR(66) NOT NULL COMMENT '交易哈希',
    collected_data JSON NOT NULL COMMENT '采集到的数据',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY `idx_task_id` (`task_id`),
    KEY `idx_block_number` (`block_number`),
    FOREIGN KEY (`task_id`) REFERENCES collector_tasks(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='采集数据表';

-- 插入示例数据
INSERT INTO collector_configs (name, chain_type, rpc_url, api_key, retry_times, retry_interval) VALUES
('ETH_Mainnet', 'eth', 'https://eth-mainnet.alchemyapi.io/v2/your-api-key', 'your-api-key', 3, 1000),
('BSC_Mainnet', 'bsc', 'https://bsc-dataseed.binance.org/', '', 3, 1000);

-- collection_fields 示例JSON结构:
-- {
--     "event_signature": "Transfer(address,address,uint256)",
--     "fields": [
--         {
--             "name": "from",
--             "type": "address",
--             "index": 0
--         },
--         {
--             "name": "to",
--             "type": "address",
--             "index": 1
--         },
--         {
--             "name": "value",
--             "type": "uint256",
--             "index": 2
--         }
--     ],
--     "filters": {
--         "address": ["0x..."],
--         "topics": ["0x..."]
--     }
-- }

INSERT INTO collector_tasks (config_id, name, description, contract_address, start_block, collection_fields, task_interval, batch_size) VALUES
(1, 'USDT_Transfer', 'USDT Transfer Events', '0xdac17f958d2ee523a2206206994597c13d831ec7', 12000000,
'{
    "event_signature": "Transfer(address,address,uint256)",
    "fields": [
        {
            "name": "from",
            "type": "address",
            "index": 0
        },
        {
            "name": "to",
            "type": "address",
            "index": 1
        },
        {
            "name": "value",
            "type": "uint256",
            "index": 2
        }
    ],
    "filters": {
        "address": ["0xdac17f958d2ee523a2206206994597c13d831ec7"],
        "topics": ["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"]
    }
}', 5000, 100);
