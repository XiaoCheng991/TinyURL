# TinyURL 技术栈研究

## 推荐技术栈

### 核心语言
- **Go 1.21+** — 高性能、静态编译、内置并发支持
- 优势：简洁的语法、优秀的标准库、原生 goroutine 协程

### Web 框架
- **Gin** — GitHub star 最多，性能优秀
  - 中间件生态丰富
  - 路由分组方便
  - JSON 处理高效

### 数据库
- **MySQL 8.0+** — 主存储
  - 适合 URL 映射这种结构化数据
  - 索引优化后查询性能优秀
  - 社区成熟，资料丰富
- **Redis 7.x** — 缓存层
  - 热点数据缓存
  - 短码自增器
  - 限流计数器

### 开发工具
- **GoLand** — IDE（学习者使用）
- **gofmt** — 代码格式化
- **golint** — 代码检查
- **testing** — 单元测试
- **Testify** — 断言库

## 各层技术选型理由

| 层级 | 技术 | 理由 |
|------|------|------|
| 语言 | Go 1.21+ | 高性能、内置并发、学习曲线适中 |
| 框架 | Gin | 生态成熟、性能优秀、文档丰富 |
| 主存 | MySQL | 成熟稳定、索引高效、SQL 查询灵活 |
| 缓存 | Redis | 读写性能高、数据结构丰富 |
| 测试 | testing + Testify | 官方标准、断言友好 |

## 依赖版本建议

```
go 1.21+
gin v1.9+
gorm v2.0+ 或 sqlx
redis/v9
mysql v1.4+
testify
```

## 不推荐的技术

- **直接使用 net/http 而不用 Gin** — Gin 生态更完整，中间件丰富
- **只用 Redis 不用 MySQL** — 数据持久化很重要
- **使用 MongoDB** — 对本项目来说过于复杂

---

*Sources:*
- [Go TinyURL short URL service architecture best practices](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)
- [Go Gin framework Redis MySQL implementation](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)