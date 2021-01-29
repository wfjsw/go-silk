package skpsilk

// silk/src/SKP_Silk_k2a_Q16.c

func k2a_Q16(A_Q24 []int32, rc_Q16 []int32, order int32) {
	var k, n int32
	Atmp := make([]int32, MAX_ORDER_LPC)
	for k = 0; k < order; k++ {
		for n = 0; n < k; n++ {
			Atmp[n] = A_Q24[n]
		}
		for n = 0; n < k; n++ {
			A_Q24[n] = SMLAWW(A_Q24[n], Atmp[k-n-1], rc_Q16[k])
		}
		A_Q24[k] = -(rc_Q16[k] << 8)
	}
}
