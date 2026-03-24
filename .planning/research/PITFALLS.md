# TinyURL 开发陷阱与注意事项

## 常见错误

### 1. 短码冲突
**问题**：随机生成短码时可能产生重复
**预防**：
- 使用自增 ID + Base62 编码
- 或随机生成后检查数据库唯一性

### 2. 缓存不一致
**问题**：MySQL 和 Redis 数据不同步
**预防**：
- 写入时双写
- 使用事务保证原子性
- 设置合理的缓存过期时间

### 3. 数据库索引缺失
**问题**：查询性能差
**预防**：
- 为 `short_code` 列建立唯一索引
- 为 `original_url` 适当建索引

### 4. 恶意请求
**问题**：暴力创建短链、爬虫访问
**预防**：
- IP 限流
- 请求频率限制
- 输入验证

## 安全风险

| 风险 | 说明 | 防护措施 |
|------|------|----------|
| 开放重定向 | 被用来做钓鱼链接 | 校验目标URL白名单 |
| URL 注入 | 恶意构造 URL | 输入验证、长度限制 |
| 枚举攻击 | 遍历短码 | 使用无序ID |
| 隐私泄露 | 点击数据泄露 | 脱敏处理、日志保护 |

## 性能问题

### 高频写入
**问题**：短时间内大量创建请求
**解决方案**：
- 批量写入
- 消息队列削峰

### 高频读取
**问题**：热门短链访问压力
**解决方案**：
- 多级缓存
- Redis 集群
- 读写分离

## 开发建议

### 学习阶段
1. 先实现核心功能，再加缓存
2. 每个功能完成后写测试
3. 先本地测试通过再进行下一步

### 常见问题排查
```bash
# 查看服务日志
tail -f logs/app.log

# 检查数据库连接
mysql -u root -p -e "SHOW PROCESSLIST"

# 检查 Redis 连接
redis-cli ping
```

## 代码审查清单

- [ ] 错误处理是否完整
- [ ] 是否有单元测试覆盖
- [ ] 输入验证是否严格
- [ ] 是否有日志记录
- [ ] 配置文件是否敏感信息分离

---

*Sources:*
- [Short URL service common problems and security concerns](https://www.reddit.com/r/shorturl/comments/)
- [Go TinyURL architecture best practices](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)