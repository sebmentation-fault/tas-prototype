package assert

// Assert whether a pointer points to something
func NotNil(p interface{}, errMsg string) {
	if p != nil {
		return
	}
	panic(errMsg)
}

// Assert whether a string is empty
func NotEmpty(s string, errMsg string) {
	if s != "" {
		return
	}
	panic(errMsg)
}
