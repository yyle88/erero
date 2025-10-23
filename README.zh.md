[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/erero/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/erero/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/erero)](https://pkg.go.dev/github.com/yyle88/erero)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/erero/main.svg)](https://coveralls.io/github/yyle88/erero?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yyle88/erero)
[![GitHub Release](https://img.shields.io/github/release/yyle88/erero.svg)](https://github.com/yyle88/erero/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/erero)](https://goreportcard.com/report/github.com/yyle88/erero)

# erero

**erero** 是一个简单的错误处理包，旨在记录错误并附带其上下文和位置。

**erero** 包名避免与 Go 的标准 `errors` 包和知名包（如 `github.com/pkg/errors` 和 `github.com/go-kratos/kratos/v2/errors`）冲突。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 安装

```bash
go get github.com/yyle88/erero
```

## 示例用法

```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/yyle88/erero"
)

func sub1() error {
	if rand.IntN(100) < 30 {
		return erero.New("wrong")
	}
	return nil
}

func sub2() error {
	if num := rand.IntN(100); num < 20 {
		return erero.Errorf("wrong num=%d", num)
	}
	return nil
}

func run() error {
	if err := sub1(); err != nil {
		return erero.WithMessage(err, "sub1 is wrong")
	}
	if err := sub2(); err != nil {
		return erero.WithMessagef(err, "sub2 is wrong")
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(erero.Wro(err))
	}
	fmt.Println("success")
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 错误包装

此示例展示如何使用 `Wrap` 和 `Wrapf` 包装错误并添加堆栈跟踪。

```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/yyle88/erero"
)

func readConfig() error {
	if rand.IntN(100) < 30 {
		return erero.New("config file not found")
	}
	return nil
}

func initDatabase() error {
	if rand.IntN(100) < 20 {
		return erero.New("connection timeout")
	}
	return nil
}

func setupApp() error {
	if err := readConfig(); err != nil {
		return erero.Wrap(err, "failed to read config")
	}
	if err := initDatabase(); err != nil {
		return erero.Wrapf(err, "failed to init database with timeout=%ds", 30)
	}
	return nil
}

func main() {
	if err := setupApp(); err != nil {
		fmt.Printf("Setup failed: %+v\n", err)
		return
	}
	fmt.Println("App setup success")
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

### 错误传播

此示例展示错误在多个函数间的传播。

```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/yyle88/erero"
)

func validateAge(age int) error {
	if age < 0 {
		return erero.Errorf("invalid age: %d", age)
	}
	if age < 18 {
		return erero.New("age too young")
	}
	return nil
}

func validateScore(score int) error {
	if score < 0 || score > 100 {
		return erero.Errorf("invalid score: %d", score)
	}
	return nil
}

func processData(age, score int) error {
	if err := validateAge(age); err != nil {
		return erero.WithMessage(err, "age validation failed")
	}
	if err := validateScore(score); err != nil {
		return erero.WithMessage(err, "score validation failed")
	}
	return nil
}

func main() {
	age := rand.IntN(25)
	score := rand.IntN(120)

	fmt.Printf("Processing: age=%d, score=%d\n", age, score)

	if err := processData(age, score); err != nil {
		fmt.Printf("Error: %v\n", erero.Wro(err))
		return
	}
	fmt.Println("Data processed success")
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
