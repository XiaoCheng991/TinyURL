# TinyURL 研究总结

## 核心发现

### 技术栈
- **Go + Gin + MySQL + Redis** 是标准技术栈
- 使用自增 ID + Base62 编码生成短码
- Cache-Aside 模式实现缓存

### 功能优先级
1. MVP：基础 CRUD + MySQL + Redis 缓存
2. 进阶：自定义别名、过期时间
3. 完善：访问统计、限流防护

### 关键陷阱
1. 短码冲突 — 使用自增 ID 预防
2. 缓存不一致 — 双写 + 过期策略
3. 安全风险 — 输入验证 + 限流

## 表格：学习路径

| 阶段 | 目标 | 关键技术 |
|------|------|----------|
| Phase 1 | 基础设施 | Go 模块、Gin、中间件 |
| Phase 2 | 核心功能 | MySQL CRUD、API 设计 |
| Phase 3 | 性能优化 | Redis 缓存、Cache-Aside |
| Phase 4 | 质量保证 | 单元测试、集成测试 |
| Phase 5 | 高级特性 | 自定义别名、统计、限流 |

## 推荐顺序

1. 先实现核心短链功能
2. 添加 Redis 缓存
3. 编写测试
4. 实现高级特性

## 学习资源

- Go 官方文档：https://go.dev/doc/
- Gin 官方文档：https://gin-gonic.com/docs/
- GORM 官方文档：https://gorm.io/docs/

---

*Research completed: 2026-03-24*