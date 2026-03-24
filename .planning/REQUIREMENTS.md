# TinyURL 需求文档

## Traceability

| REQ-ID | 需求 | 阶段 | 状态 |
|--------|------|------|------|
| CORE-01 | 创建短链 | Phase 2 | 待实现 |
| CORE-02 | 短链跳转 | Phase 2 | 待实现 |
| CORE-03 | 短链查询 | Phase 2 | 待实现 |
| DATA-01 | MySQL 持久化 | Phase 2 | 待实现 |
| CACHE-01 | Redis 缓存 | Phase 3 | 待实现 |
| TEST-01 | 单元测试 | Phase 4 | 待实现 |
| TEST-02 | 集成测试 | Phase 4 | 待实现 |
| ADV-01 | 自定义别名 | Phase 5 | 待实现 |
| ADV-02 | 过期时间 | Phase 5 | 待实现 |
| ADV-03 | 访问统计 | Phase 5 | 待实现 |
| SEC-01 | 限流防护 | Phase 5 | 待实现 |

## v1 Requirements

### CORE — 核心功能

- [ ] **CORE-01**: 长链转短链 — 用户可以提交长URL，获得对应的短URL
- [ ] **CORE-02**: 短链跳转 — 访问短URL时，重定向到原始长URL（302状态码）
- [ ] **CORE-03**: 短链查询 — 根据短码查询对应的长URL

### DATA — 数据存储

- [ ] **DATA-01**: MySQL持久化 — URL映射关系存储到MySQL数据库
- [ ] **DATA-02**: 数据表设计 — 设计url_mapping表结构

### CACHE — 缓存

- [ ] **CACHE-01**: Redis缓存 — 热点短链数据缓存到Redis
- [ ] **CACHE-02**: 缓存更新 — 写入时同步更新缓存

### TEST — 测试

- [ ] **TEST-01**: 单元测试 — 核心函数有单元测试覆盖
- [ ] **TEST-02**: 集成测试 — API接口有集成测试覆盖

## v2 Requirements（进阶）

### ADV — 高级特性

- [ ] **ADV-01**: 自定义别名 — 用户可以指定自定义短码
- [ ] **ADV-02**: 过期时间 — 支持设置短链的有效期
- [ ] **ADV-03**: 访问统计 — 记录并展示短链的访问次数

### SEC — 安全

- [ ] **SEC-01**: 限流防护 — 防止恶意暴力创建短链
- [ ] **SEC-02**: 输入验证 — 校验URL格式和长度

## Out of Scope

- **用户系统** — MVP阶段不需要，后续项目再实现
- **JWT认证** — API认证复杂度高，留待后续
- **Docker部署** — 本地学习不需要容器化
- **分布式部署** — 单机版本足够学习

---

*Requirements defined: 2026-03-24*