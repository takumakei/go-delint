package delint

// AnywayFunc calls fn, discards its result.
// It should be used in situations where lint reports that the result is
// unchecked, but the result of function is irrelevant.
//
// In the situation that a result is irrelevant the blank identifier is used in
// general whereas it cannot be called directly with defer.
// (https://golang.org/doc/effective_go#blank)
//
// example
//
//     // LGTM, the way with AnywayFunc
//     defer delint.AnywayFunc(hello)
//
//     // Okay, the way with the blank identifier
//     defer func() {
//     	_ = hello()
//     }()
func AnywayFunc(fn func() error) {
	_ = fn()
}
