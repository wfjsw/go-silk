package skpsilk

// silk/src/SKP_Silk_inlines.go

/* count leading zeros of SKP_int64 */
func CLZ64(in int64) int32 {
	in_upper := int32(in >> 32)
	if in_upper == 0 {
		return 32 + CLZ32(int32(in))
	} else {
		/* Search in the upper 32 bits */
		return CLZ32(in_upper)
	}
}

/* get number of leading zeros and fractional part (the bits right after the leading one */
func CLZ_FRAC(in int32) (lz int32, frac_Q7 int32) {
	lzeros := CLZ32(in)
	return lzeros, ROR32(in, int(24-lzeros)) & 0x7f
}

/* Approximation of square root                                          */
/* Accuracy: < +/- 10%  for output values > 15                           */
/*           < +/- 2.5% for output values > 120                          */

func SQRT_APPROX(x int32) int32 {
	if x <= 0 {
		return 0
	}

	lz, frac_Q7 := CLZ_FRAC(x)

	var y int32

	if lz&1 != 0 {
		y = 32768
	} else {
		y = 46214 /* 46214 = sqrt(2) * 32768 */
	}

	/* get scaling right */
	y >>= lz >> 1

	y = SMLAWB(y, y, SMULBB(213, frac_Q7))

	return y
}
