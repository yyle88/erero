[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/erero/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/erero/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/erero)](https://pkg.go.dev/github.com/yyle88/erero)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/erero/master.svg)](https://coveralls.io/github/yyle88/erero?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/erero.svg)](https://github.com/yyle88/erero/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/erero)](https://goreportcard.com/report/github.com/yyle88/erero)

# erero

**erero** is a simple error-handling package designed to log errors along with their context and location.

package name **erero** doesn't conflict with Go's standard `errors` package or other popular libraries like `github.com/pkg/errors` or `github.com/go-kratos/kratos/v2/errors`.

## CHINESE README

[中文说明](README.zh.md)

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

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/erero.svg?variant=adaptive)](https://starchart.cc/yyle88/erero)
