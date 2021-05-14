package delint_test

import (
	"fmt"

	"github.com/takumakei/go-delint"
)

func ExampleMust_doNotDoThis() {
	fmt.Println("1st")

	hello := func() error {
		fmt.Println("2nd, not 3rd")
		return nil
	}
	// DO NOT DO THIS
	// THIS MUST BE DIFFERENT FROM YOUR INTENTION
	// hello() is called immediately, Must is called using its result later
	// before exiting the function.
	defer delint.Must(hello())

	fmt.Println("3rd")
	// Output:
	// 1st
	// 2nd, not 3rd
	// 3rd
}
