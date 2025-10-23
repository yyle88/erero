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
