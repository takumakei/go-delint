package delint_test

import (
	"errors"
	"fmt"

	"github.com/takumakei/go-delint"
)

func ExampleAnywayFunc() {
	hello := func() error {
		fmt.Println("3rd")
		return errors.New("")
	}
	fmt.Println("1st")
	defer delint.AnywayFunc(hello)
	fmt.Println("2nd")
	// Output:
	// 1st
	// 2nd
	// 3rd
}
