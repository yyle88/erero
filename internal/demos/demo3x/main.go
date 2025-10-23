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
