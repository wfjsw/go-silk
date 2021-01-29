package skpsilk

// silk/src/SKP_Silk_k2a.c

func k2a(A_Q24 []int32, rc_Q15 []int16, order int32) {
	var k, n int32
	Atmp := make([]int32, MAX_ORDER_LPC)
	for k = 0; k < order; k++ {
		for n = 0; n < k; n++ {
			Atmp[n] = A_Q24[n]
		}
		for n = 0; n < k; n++ {
			A_Q24[n] = SMLAWB(A_Q24[n], Atmp[k-n-1]<<1, int32(rc_Q15[k]))
		}
		A_Q24[k] = -(int32(rc_Q15[k]) << 9)
	}
}
