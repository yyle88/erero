# erero

**erero** 是一个简单的错误处理包，旨在记录错误并附带其上下文和位置。

**erero** 包名避免与 Go 的标准 `errors` 包或其他流行的库（如 `github.com/pkg/errors` 或 `github.com/go-kratos/kratos/v2/errors`）冲突。

## 英文文档

[English README](README.md)

## 安装

```bash
go get github.com/yyle88/erero
```

## 示例用法

```go
package main

import (
	"fmt"
	"math/rand"

	"github.com/yyle88/erero"
)

func sub1() error {
	if rand.Intn(100) < 30 {
		return erero.New("wrong")
	}
	return nil
}

func sub2() error {
	if num := rand.Intn(100); num < 20 {
		return erero.Errorf("wrong num=%d", num)
	}
	return nil
}

func run() error {
	if err := sub1(); err != nil {
		return erero.WithMessage(err, "sub1 发生错误")
	}
	if err := sub2(); err != nil {
		return erero.WithMessagef(err, "sub2 发生错误")
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(erero.Wro(err))
	}
	fmt.Println("成功")
}
```

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
