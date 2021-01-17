package skpsilk

// silk/src/SKP_Silk_bwexpander.c

func bwexpander(ar []int32, d int, chirp_Q16 int32) {
	chirp_minus_one_Q16 := chirp_Q16 - 65536

	/* NB: Dont use SKP_SMULWB, instead of SKP_RSHIFT_ROUND( SKP_MUL() , 16 ), below. */
	/* Bias in SKP_SMULWB can lead to unstable filters                                */
	for i := 0; i < d-1; i++ {
		ar[i] = RSHIFT_ROUND(chirp_Q16*ar[i], 16) & 0xFFFF
		chirp_Q16 += RSHIFT_ROUND(chirp_Q16*chirp_minus_one_Q16, 16)
	}
	ar[d-1] = RSHIFT_ROUND(chirp_Q16*ar[d-1], 16)
}
