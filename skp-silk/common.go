package skpsilk

func assert(cond bool) {
	if !cond {
		panic("Assertion failed")
	}
}
