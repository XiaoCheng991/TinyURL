# Go 语言学习目录

> TinyURL 项目配套学习资料

## 快速导航

### 基础概念
| 文档 | 内容 | 状态 |
|------|------|------|
| [01-变量与类型.md](./01-变量与类型.md) | 变量声明、指针、值传递 | ✅ |
| [02-控制流.md](./02-控制流.md) | for、if、switch | ✅ (合并到01) |
| [03-函数.md](./03-函数.md) | 多返回值、命名返回、可变参数 | ✅ (合并到01) |
| [04-数据结构.md](./04-数据结构.md) | Slice、Map、String | ✅ (合并到01) |
| [05-结构体与接口.md](./05-结构体与接口.md) | struct、method、interface | ✅ (合并到01) |
| [06-错误处理.md](./06-错误处理.md) | error、多返回值、errors 包 | ✅ (合并到01) |
| [07-defer.md](./07-defer.md) | defer 执行时机、常见用法 | ✅ (合并到01) |
| [08-并发编程.md](./08-并发编程.md) | Goroutine、Channel、锁 | ✅ (合并到01) |

### 进阶主题
| 文档 | 内容 |
|------|------|
| [10-泛型.md](./10-泛型.md) | Go 1.18+ 泛型 |
| [11-反射.md](./11-反射.md) | reflect 包 |
| [12-测试.md](./12-测试.md) | 单元测试、基准测试 |
| [13-性能优化.md](./13-性能优化.md) | 内存分配、逃逸分析 |

### Gin 框架专题
| 文档 | 内容 |
|------|------|
| [20-Gin入门.md](./20-Gin入门.md) | Gin 基础、路由、中间件 |
| [21-Gin进阶.md](./21-Gin进阶.md) | 参数绑定、验证器、分组路由 |
| [22-Gin最佳实践.md](./22-Gin最佳实践.md) | 项目结构、中间件设计 |

## 学习路径

```
第1天 → Go 基础语法（01-07）
第2天 → 并发编程（08）
第3天 → Gin 框架入门（20）
第4天 → 开始 TinyURL 项目实战
```

## 配套代码

```
docs/go-basics/
├── examples/
│   ├── variables.go       # 变量与指针示例
│   ├── control_flow.go    # 控制流示例
│   ├── functions.go       # 函数示例
│   ├── slice_map.go       # Slice 和 Map 示例
│   ├── struct_interface.go # 结构体和接口示例
│   ├── error_handling.go  # 错误处理示例
│   ├── defer_demo.go      # defer 示例
│   ├── goroutine_demo.go  # 并发示例
│   └── gin_demo.go        # Gin 示例
└── quiz/
    └── answers.md         # 面试题答案详解
```

## 运行示例代码

```bash
cd docs/go-basics/examples
go run variables.go
```

## 面试题速查

| 难度 | 必问题目 |
|------|----------|
| ⭐ | Slice vs Array 的区别 |
| ⭐ | make vs new 的区别 |
| ⭐⭐ | Map 是线程安全的吗？ |
| ⭐⭐ | channel 关闭后还能读取吗？ |
| ⭐⭐⭐ | GMP 调度模型详细解释 |
| ⭐⭐⭐ | 什么是逃逸分析？ |
| ⭐⭐⭐ | select 的实现原理？ |

---

*持续更新中...*