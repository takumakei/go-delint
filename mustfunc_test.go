package delint_test

import (
	"fmt"

	"github.com/takumakei/go-delint"
)

func ExampleMustFunc() {
	hello := func() error {
		fmt.Println("3rd")
		return nil
	}
	fmt.Println("1st")
	defer delint.MustFunc(hello)
	fmt.Println("2nd")
	// Output:
	// 1st
	// 2nd
	// 3rd
}
