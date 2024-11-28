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

## 许可协议

此项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE) 文件。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉
