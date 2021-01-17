package skpsilk

// silk/src/SKP_Silk_interpolate.c

// interpolate Interpolate two vectors
func interpolate(xi []int, x0 []int, x1 []int, ifact_Q2, d int) {
	assert(ifact_Q2 >= 0)
	assert(ifact_Q2 <= (1 << 2))

	for i := 0; i < d; i++ {
		xi[i] = int(int32(x0[i]) + (int32(x1[i])-int32(x0[i])*int32(ifact_Q2))>>2)
	}
}
