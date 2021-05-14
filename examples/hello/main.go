package main

func main() {
	// Some lint will report that result of the function is unchecked,
	// even though the author knew the function never returns err!=nil.
	hello()
}

func hello() error {
	return nil
}
