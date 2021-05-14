package delint

// MustFunc calls Must(fn()).
//
// This is intended to use with defer.
//
//     // LGTM. the way with MustFunc
//     defer delint.MustFunc(hello)
//
//     // Okay, the way without MustFunc
//     defer func() {
//     	delint.Must(hello())
//     }
//
//     // DO NOT DO THIS
//     // THIS MUST BE DIFFERENT FROM YOUR INTENTION
//     // hello() is called immediately, Must is called using its result later
//     // before exiting the function.
//     defer Must(hello())
func MustFunc(fn func() error) {
	Must(fn())
}
