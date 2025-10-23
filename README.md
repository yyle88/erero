[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/erero/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/erero/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/erero)](https://pkg.go.dev/github.com/yyle88/erero)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/erero/main.svg)](https://coveralls.io/github/yyle88/erero?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yyle88/erero)
[![GitHub Release](https://img.shields.io/github/release/yyle88/erero.svg)](https://github.com/yyle88/erero/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/erero)](https://goreportcard.com/report/github.com/yyle88/erero)

# erero

**erero** is a simple err handling package designed to log errors with context and location.

package name **erero** doesn't conflict with Go's standard `errors` package and famous packages like `github.com/pkg/errors` and `github.com/go-kratos/kratos/v2/errors`.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

```bash
go get github.com/yyle88/erero
```

## Example Usage

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

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Error Wrapping

This example shows how to wrap errors with stack traces using `Wrap` and `Wrapf`.

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

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### Error Propagation

This example demonstrates error propagation across multiple functions.

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

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->
