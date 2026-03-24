# GSD 命令参考

**GSD** (Get Shit Done) 创建专为 Claude Code 单人智能开发优化的分层项目计划。

## 快速开始

1. `/gsd:new-project` - 初始化项目（包括研究、需求、路线图）
2. `/gsd:plan-phase 1` - 为第一阶段创建详细计划
3. `/gsd:execute-phase 1` - 执行该阶段

## 保持更新

GSD 发展迅速。定期更新：

```bash
npx get-shit-done-cc@latest
```

## 核心工作流程

```
/gsd:new-project → /gsd:plan-phase → /gsd:execute-phase → 重复
```

### 项目初始化

**`/gsd:new-project`**
通过统一流程初始化新项目。

一个命令即可从想法进入准备规划状态：
- 深度提问以了解你要构建什么
- 可选的领域研究（启动 4 个并行研究智能体）
- 定义包含 v1/v2/范围外划分的需求
- 创建包含阶段划分和成功标准的路线图

创建所有 `.planning/` 工件：
- `PROJECT.md` — 项目愿景
- `config.json` — 工作流程模式（interactive/yolo）
- `research/` — 领域研究（如果选择）
- `REQUIREMENTS.md` — 带有 REQ-ID 的范围化需求
- `ROADMAP.md` — 映射到需求的阶段
- `STATE.md` — 项目记忆

用法：`/gsd:new-project`

**`/gsd:map-codebase`**
为现有代码库创建映射（用于既有项目）。

- 使用并行 Explore 智能体分析代码库
- 创建包含 7 个专注文档的 `.planning/codebase/`
- 涵盖技术栈、架构、结构、约定、测试、集成、关注点
- 在既有项目的 `/gsd:new-project` 之前使用

用法：`/gsd:map-codebase`

### 阶段规划

**`/gsd:discuss-phase <number>`**
在规划之前帮助你阐述对某个阶段的愿景。

- 捕获你对这个阶段如何运作的想象
- 创建包含你的愿景、核心要素和边界的 CONTEXT.md
- 当你对某个东西应该是什么样子/感觉有想法时使用
- 可选 `--batch` 一次性询问 2-5 个相关问题，而不是一个接一个

用法：`/gsd:discuss-phase 2`
用法：`/gsd:discuss-phase 2 --batch`
用法：`/gsd:discuss-phase 2 --batch=3`

**`/gsd:research-phase <number>`**
针对小众/复杂领域的综合生态研究。

- 发现标准技术栈、架构模式、陷阱
- 创建包含"专家如何构建这个"知识的 RESEARCH.md
- 用于 3D、游戏、音频、着色器、机器学习和其他专业领域
- 超越"使用哪个库"到生态系统知识

用法：`/gsd:research-phase 3`

**`/gsd:list-phase-assumptions <number>`**
查看 Claude 在开始之前计划做什么。

- 显示 Claude 对某个阶段的预期方法
- 如果 Claude 误解了你的愿景，让你重新调整
- 不创建文件 — 仅对话输出

用法：`/gsd:list-phase-assumptions 3`

**`/gsd:plan-phase <number>`**
为特定阶段创建详细的执行计划。

- 生成 `.planning/phases/XX-phase-name/XX-YY-PLAN.md`
- 将阶段分解为具体的、可操作的任务
- 包括验证标准和成功措施
- 每个阶段支持多个计划（XX-01、XX-02 等）

用法：`/gsd:plan-phase 1`
结果：创建 `.planning/phases/01-foundation/01-01-PLAN.md`

**PRD 快捷路径：** 传递 `--prd path/to/requirements.md` 以完全跳过 discuss-phase。你的 PRD 成为 CONTEXT.md 中的锁定决策。

### 执行

**`/gsd:execute-phase <phase-number>`**
执行阶段中的所有计划，或运行特定波次。

- 按波次分组计划（来自 frontmatter），按顺序执行波次
- 每个波次内的计划通过 Task 工具并行运行
- 可选的 `--wave N` 标志仅执行波次 `N`，除非阶段现已完全完成，否则停止
- 所有计划完成后验证阶段目标
- 更新 REQUIREMENTS.md、ROADMAP.md、STATE.md

用法：`/gsd:execute-phase 5`
用法：`/gsd:execute-phase 5 --wave 2`

### 智能路由器

**`/gsd:do <description>`**
自动将自由文本路由到正确的 GSD 命令。

- 分析自然语言输入以找到最佳匹配的 GSD 命令
- 充当调度器 — 永不自已完成工作
- 通过要求你在顶级匹配之间进行选择来解决歧义

用法：`/gsd:do 修复登录按钮`
用法：`/gsd:do 重构认证系统`
用法：`/gsd:do 我想开始一个新的里程碑`

### 快速模式

**`/gsd:quick [--full] [--discuss] [--research]`**
使用 GSD 保证执行小的、临时任务，但跳过可选的智能体。

快速模式使用相同的系统但路径更短：
- 启动规划器 + 执行器（默认跳过研究器、检查器、验证器）
- 快速任务位于独立于规划阶段的 `.planning/quick/`
- 更新 STATE.md 跟踪（而非 ROADMAP.md）

标志启用额外的质量步骤：
- `--discuss` — 轻量级讨论以在规划前暴露灰色地带
- `--research` — 专注的研究智能体在规划前调查方法
- `--full` — 添加计划检查（最多 2 次迭代）和执行后验证

用法：`/gsd:quick`
用法：`/gsd:quick --research --full`

---

**`/gsd:fast [description]`**
内联执行琐碎任务 — 无子智能体，无规划文件，无开销。

用于太琐碎而无法证明规划合理的任务：拼写错误修复、配置更改、遗忘的提交、简单添加。

用法：`/gsd:fast "修复 README 中的拼写错误"`
用法：`/gsd:fast "将 .env 添加到 gitignore"`

### 路线图管理

**`/gsd:add-phase <description>`**
将新阶段添加到当前里程碑的末尾。

用法：`/gsd:add-phase "添加管理面板"`

**`/gsd:insert-phase <after> <description>`**
在现有阶段之间插入紧急工作作为小数阶段。

用法：`/gsd:insert-phase 7 "修复关键认证漏洞"`

**`/gsd:remove-phase <number>`**
移除未来阶段并重新编号后续阶段。

用法：`/gsd:remove-phase 17`

### 里程碑管理

**`/gsd:new-milestone <name>`**
通过统一流程开始新里程碑。

用法：`/gsd:new-milestone "v2.0 功能"`
用法：`/gsd:new-milestone --reset-phase-numbers "v2.0 功能"`

**`/gsd:complete-milestone <version>`**
存档完成的里程碑并为下一个版本做准备。

用法：`/gsd:complete-milestone 1.0.0`

### 进度跟踪

**`/gsd:progress`**
检查项目状态并智能路由到下一个操作。

用法：`/gsd:progress`

### 会话管理

**`/gsd:resume-work`**
使用完整上下文恢复从中断的会话中进行工作。

用法：`/gsd:resume-work`

**`/gsd:pause-work`**
在阶段中途暂停工作时创建上下文交接。

用法：`/gsd:pause-work`

### 调试

**`/gsd:debug [issue description]`**
通过跨上下文重置的持久状态进行系统性调试。

用法：`/gsd:debug "登录按钮不工作"`
用法：`/gsd:debug`（恢复活动会话）

### 快速笔记

**`/gsd:note <text>`**
零摩擦想法捕获 — 一个命令，即时保存。

用法：`/gsd:note 重构钩子系统`
用法：`/gsd:note list`
用法：`/gsd:note promote 3`
用法：`/gsd:note --global 跨项目想法`

### 待办事项管理

**`/gsd:add-todo [description]`**
从当前对话捕获想法或任务作为待办事项。

用法：`/gsd:add-todo`（从对话推断）
用法：`/gsd:add-todo 添加认证令牌刷新`

**`/gsd:check-todos [area]`**
列出待处理待办事项并选择一个来处理。

用法：`/gsd:check-todos`
用法：`/gsd:check-todos api`

### 用户验收测试

**`/gsd:verify-work [phase]`**
通过对话式 UAT 验证构建的功能。

用法：`/gsd:verify-work 3`

### 交付工作

**`/gsd:ship [phase]`**
从完成的阶段工作中创建带有自动生成的正文的 PR。

用法：`/gsd:ship 4` 或 `/gsd:ship 4 --draft`

---

**`/gsd:review --phase N [--gemini] [--claude] [--codex] [--all]`**
跨 AI 同行审查 — 调用外部 AI CLI 独立审查阶段计划。

用法：`/gsd:review --phase 3 --all`

---

**`/gsd:pr-branch [target]`**
通过过滤掉 .planning/ 提交为拉取请求创建干净分支。

用法：`/gsd:pr-branch` 或 `/gsd:pr-branch main`

---

**`/gsd:plant-seed [idea]`**
使用触发条件捕获前瞻性想法以自动浮现。

用法：`/gsd:plant-seed "在我们构建事件系统时添加实时通知"`

---

**`/gsd:audit-uat`**
所有未决 UAT 和验证项目的跨阶段审计。

用法：`/gsd:audit-uat`

### 里程碑审计

**`/gsd:audit-milestone [version]`**
根据原始意图审计里程碑完成情况。

用法：`/gsd:audit-milestone`

**`/gsd:plan-milestone-gaps`**
创建阶段以关闭审计识别的差距。

用法：`/gsd:plan-milestone-gaps`

### 配置

**`/gsd:settings`**
以交互方式配置工作流程开关和模型配置文件。

用法：`/gsd:settings`

**`/gsd:set-profile <profile>`**
为 GSD 智能体快速切换模型配置文件。

- `quality` — Opus 到处使用，除了验证
- `balanced` — Opus 用于规划，Sonnet 用于执行（默认）
- `budget` — Sonnet 用于编写，Haiku 用于研究/验证
- `inherit` — 为所有智能体使用当前会话模型

用法：`/gsd:set-profile budget`

### 实用命令

**`/gsd:cleanup`**
存档来自已完成里程碑的累积阶段目录。

用法：`/gsd:cleanup`

**`/gsd:help`**
显示此命令参考。

**`/gsd:update`**
更新 GSD 到最新版本并预览更改日志。

用法：`/gsd:update`

**`/gsd:join-discord`**
加入 GSD Discord 社区。

用法：`/gsd:join-discord`

## 文件和结构

```
.planning/
├── PROJECT.md            # 项目愿景
├── ROADMAP.md            # 当前阶段划分
├── STATE.md              # 项目记忆和上下文
├── RETROSPECTIVE.md      # 活跃回顾（每次里程碑更新）
├── config.json           # 工作流程模式和门槛
├── todos/                # 捕获的想法和任务
│   ├── pending/          # 等待处理的待办事项
│   └── done/             # 已完成的待办事项
├── debug/                # 活跃调试会话
│   └── resolved/         # 存档的已解决问题
├── milestones/
│   ├── v1.0-ROADMAP.md       # 存档的路线图快照
│   ├── v1.0-REQUIREMENTS.md  # 存档的需求
│   └── v1.0-phases/          # 存档的阶段目录
├── codebase/             # 代码库映射
│   ├── STACK.md          # 语言、框架、依赖
│   ├── ARCHITECTURE.md   # 模式、层、数据流
│   ├── STRUCTURE.md      # 目录布局、关键文件
│   ├── CONVENTIONS.md    # 编码标准、命名
│   ├── TESTING.md        # 测试设置、模式
│   ├── INTEGRATIONS.md   # 外部服务、API
│   └── CONCERNS.md       # 技术债务、已知问题
└── phases/
    ├── 01-foundation/
    │   ├── 01-01-PLAN.md
    │   └── 01-01-SUMMARY.md
    └── 02-core-features/
        ├── 02-01-PLAN.md
        └── 02-01-SUMMARY.md
```

## 工作流程模式

在 `/gsd:new-project` 期间设置：

**交互式模式**
- 确认每个主要决策
- 在检查点暂停以进行批准

**YOLO 模式**
- 自动批准大多数决策
- 执行计划无需确认

## 规划配置

在 `.planning/config.json` 中配置如何管理规划工件：

**`planning.commit_docs`**（默认：`true`）
- `true`：规划工件已提交到 git（标准工作流程）
- `false`：规划工件仅本地保存，不提交

## 常见工作流程

**启动新项目：**
```
/gsd:new-project
/clear
/gsd:plan-phase 1
/clear
/gsd:execute-phase 1
```

**休息后恢复工作：**
```
/gsd:progress
```

**完成里程碑：**
```
/gsd:complete-milestone 1.0.0
/clear
/gsd:new-milestone
```
