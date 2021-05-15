Package delint implements functions for better life with lint.
======================================================================

[![GoDoc](https://pkg.go.dev/badge/github.com/takumakei/go-delint)](https://godoc.org/github.com/takumakei/go-delint)

The function 'Must'
----------------------------------------------------------------------

'Must' causes panic if err is not nil.
It should be used in situations where lint reports that the result is
unchecked, but the function always return nil for now.

Consider the function that returns of type error.

    func hello() error {
    	return nil
    }

And calling it like below.

    func main() {
    	hello()
    }

Some lint will report that result of the function is unchecked,
even though the author knew the function never returns err!=nil.

    main.go:6:7: Error return value is not checked (errcheck)
            hello()
                 ^

The function returns nil for now.
Provably it wouldn't change.
But who knows it does.

In the situations like above, using 'Must' would be more reasonable than using
the blank identifier.  https://golang.org/doc/effective_go#blank

'Must' can show that always returning nil is the specification of a function.
