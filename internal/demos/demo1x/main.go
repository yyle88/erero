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
