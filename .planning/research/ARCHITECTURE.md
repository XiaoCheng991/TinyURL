# TinyURL 架构设计研究

## 整体架构

```
┌─────────────────────────────────────────────────────────────┐
│                        API Gateway (Gin)                     │
│                    GET /{code}  |  POST /api/shorten         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        Gateway Layer                         │
│                  HTTP请求解析、参数校验、中间件               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        Logic Layer                           │
│              短码生成、URL验证、业务规则处理                  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        Repo Layer                            │
│                  MySQL CRUD、Redis 缓存操作                  │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              ▼                               ▼
┌─────────────────────┐           ┌─────────────────────┐
│      MySQL          │           │      Redis          │
│   URL Mapping       │           │   Cache + Counter   │
└─────────────────────┘           └─────────────────────┘
```

## 数据流向

### 创建短链流程
```
POST /api/shorten
  → Gateway: 参数校验
  → Logic: 生成短码
  → Repo: MySQL 插入 + Redis 缓存
  → Response: 返回短链URL
```

### 访问短链流程
```
GET /{code}
  → Gateway: 路由匹配
  → Repo: Redis 查找（未命中则查 MySQL）
  → Logic: 访问计数
  → Response: 302 重定向
```

## 短码生成算法

### 方案一：Base62 自增ID
```go
// 使用 Redis INCR 生成自增 ID
id := redis.Incr("url:id")
code := base62.Encode(id)
```

### 方案二：随机字符串 + 冲突检测
```go
// 生成随机 6 位短码
code := random.String(6)
// 检查是否已存在
if exists(code) {
    retry()
}
```

**推荐**：方案一，性能更好，无冲突风险

## 缓存策略

### Cache-Aside 模式
1. 读取：先查 Redis，未命中查 MySQL，回填 Redis
2. 写入：同时更新 MySQL 和 Redis
3. 删除：同时删除 MySQL 和 Redis

### 缓存击穿防护
- 使用分布式锁
- 设置合理的过期时间

---

*Sources:*
- [Go TinyURL short URL service architecture best practices](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)
- [Go Gin framework Redis MySQL implementation](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)