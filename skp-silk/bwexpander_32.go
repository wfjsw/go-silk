package skpsilk

// silk/src/SKP_Silk_bwexpander_32.c

func bwexpander_32(ar []int32, d int, chirp_Q16 int32) {
	tmp_chirp_Q16 := chirp_Q16
	for i := 0; i < d-1; i++ {
		ar[i] = SMULWW(ar[i], tmp_chirp_Q16)
		tmp_chirp_Q16 = SMULWW(chirp_Q16, tmp_chirp_Q16)
	}
	ar[d-1] = SMULWW(ar[d-1], tmp_chirp_Q16)
}
