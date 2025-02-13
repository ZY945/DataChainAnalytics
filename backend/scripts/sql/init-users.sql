-- 创建数据库用户
-- 采集节点用户
CREATE USER IF NOT EXISTS 'collector_user'@'%' IDENTIFIED BY 'collector_password_123';
CREATE USER IF NOT EXISTS 'collector_user'@'localhost' IDENTIFIED BY 'collector_password_123';

-- Master节点用户
CREATE USER IF NOT EXISTS 'master_user'@'%' IDENTIFIED BY 'master_password_123';
CREATE USER IF NOT EXISTS 'master_user'@'localhost' IDENTIFIED BY 'master_password_123';

-- 只读用户(用于查询)
CREATE USER IF NOT EXISTS 'reader_user'@'%' IDENTIFIED BY 'reader_password_123';
CREATE USER IF NOT EXISTS 'reader_user'@'localhost' IDENTIFIED BY 'reader_password_123';

-- 授权collector数据库权限
GRANT SELECT, INSERT, UPDATE ON collector_db.* TO 'collector_user'@'%';
GRANT SELECT, INSERT, UPDATE ON collector_db.* TO 'collector_user'@'localhost';

-- 授权master数据库权限
GRANT ALL PRIVILEGES ON blockchain_master.* TO 'master_user'@'%';
GRANT ALL PRIVILEGES ON blockchain_master.* TO 'master_user'@'localhost';

-- 授权只读权限
GRANT SELECT ON collector_db.* TO 'reader_user'@'%';
GRANT SELECT ON collector_db.* TO 'reader_user'@'localhost';
GRANT SELECT ON blockchain_master.* TO 'reader_user'@'%';
GRANT SELECT ON blockchain_master.* TO 'reader_user'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;

-- 查看用户权限
SHOW GRANTS FOR 'collector_user'@'localhost';
SHOW GRANTS FOR 'master_user'@'localhost';
SHOW GRANTS FOR 'reader_user'@'localhost';

-- 创建存储过程和函数的权限（如果需要）
-- GRANT CREATE ROUTINE ON collector_db.* TO 'collector_user'@'%';
-- GRANT CREATE ROUTINE ON blockchain_master.* TO 'master_user'@'%';

-- 安全建议：
-- 1. 在生产环境中使用更强的密码
-- 2. 限制用户访问的IP地址范围
-- 3. 定期更改密码
-- 4. 监控用户访问日志 