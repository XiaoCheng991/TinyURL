<!-- GSD:project-start source:PROJECT.md -->
## Project

**Go语言全栈工程师成长计划 - TinyURL项目**

一个为期一个月的Go语言实战学习项目，通过构建高性能短域名服务，帮助有Java基础的开发者从Go语法入门进阶到企业级分布式架构设计。项目采用渐进式开发模式，从MVP到高级特性逐步完善，每个阶段包含详细的技术文档、单元测试和代码规范示例。

**Core Value:** 通过"练中学、学中练"的实战方式，让学习者在一个月内掌握Go语言核心语法、并发编程、工程化实践和分布式架构，能够独立开发高质量的项目。

### Constraints

- **时间限制**: 一个月完成项目 —— 需要分阶段渐进式开发，先MVP再逐步增强
- **运行环境**: 本地GoLand IDE运行 —— 先不考虑生产环境部署
- **Web框架**: 必须使用Gin框架 —— Gin生态成熟，适合学习
- **错误处理**: 结构化错误 — 自定义错误类型，规范错误信息
- **代码规范**: gofmt + golint — 遵循Go官方代码风格
- **文档标准**: 最详细级别 — 代码注释、README、API文档、设计文档、代码审查清单、最佳实践说明、常见问题解答
<!-- GSD:project-end -->

<!-- GSD:stack-start source:research/STACK.md -->
## Technology Stack

## 推荐技术栈
### 核心语言
- **Go 1.21+** — 高性能、静态编译、内置并发支持
- 优势：简洁的语法、优秀的标准库、原生 goroutine 协程
### Web 框架
- **Gin** — GitHub star 最多，性能优秀
### 数据库
- **MySQL 8.0+** — 主存储
- **Redis 7.x** — 缓存层
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
## 不推荐的技术
- **直接使用 net/http 而不用 Gin** — Gin 生态更完整，中间件丰富
- **只用 Redis 不用 MySQL** — 数据持久化很重要
- **使用 MongoDB** — 对本项目来说过于复杂
- [Go TinyURL short URL service architecture best practices](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)
- [Go Gin framework Redis MySQL implementation](https://www.reddit.com/r/golang/comments/1h7kj9y/go_tinyurl_short_url_service_architecture/)
<!-- GSD:stack-end -->

<!-- GSD:conventions-start source:CONVENTIONS.md -->
## Conventions

Conventions not yet established. Will populate as patterns emerge during development.
<!-- GSD:conventions-end -->

<!-- GSD:architecture-start source:ARCHITECTURE.md -->
## Architecture

Architecture not yet mapped. Follow existing patterns found in the codebase.
<!-- GSD:architecture-end -->

<!-- GSD:workflow-start source:GSD defaults -->
## GSD Workflow Enforcement

Before using Edit, Write, or other file-changing tools, start work through a GSD command so planning artifacts and execution context stay in sync.

Use these entry points:
- `/gsd:quick` for small fixes, doc updates, and ad-hoc tasks
- `/gsd:debug` for investigation and bug fixing
- `/gsd:execute-phase` for planned phase work

Do not make direct repo edits outside a GSD workflow unless the user explicitly asks to bypass it.
<!-- GSD:workflow-end -->

## gstack

使用 gstack 的 `/browse` skill 进行所有网页浏览任务，**不要使用** `mcp__claude-in-chrome__*` 工具。

### 可用的 gstack skills:

| Skill | 用途 |
|-------|------|
| `/browse` | 网页浏览（优先于 MCP 工具） |
| `/connect-chrome` | 连接 Chrome 浏览器 |
| `/setup-browser-cookies` | 设置浏览器 cookies |
| `/office-hours` | 团队同步和答疑 |
| `/plan-ceo-review` | 规划 CEO 汇报 |
| `/plan-eng-review` | 规划工程评审 |
| `/plan-design-review` | 规划设计评审 |
| `/design-consultation` | 获取设计咨询 |
| `/design-shotgun` | 快速设计反馈 |
| `/design-review` | 设计评审 |
| `/review` | 代码评审 |
| `/ship` | 发布代码 |
| `/land-and-deploy` | 提交并部署 |
| `/canary` | 金丝雀部署 |
| `/benchmark` | 运行基准测试 |
| `/qa` | 质量保证 |
| `/qa-only` | 仅 QA 模式 |
| `/retro` | 运行回顾会 |
| `/investigate` | 调查问题 |
| `/document-release` | 文档发布 |
| `/codex` | Codex 操作 |
| `/cso` | CSO 操作 |
| `/autoplan` | 自动规划 |
| `/careful` | 谨慎模式 |
| `/freeze` | 冻结变更 |
| `/guard` | 守护模式 |
| `/unfreeze` | 解冻变更 |
| `/gstack-upgrade` | 升级 gstack |
| `/setup-deploy` | 设置部署 |



<!-- GSD:profile-start -->
## Developer Profile

> Profile not yet configured. Run `/gsd:profile-user` to generate your developer profile.
> This section is managed by `generate-claude-profile` -- do not edit manually.
<!-- GSD:profile-end -->
