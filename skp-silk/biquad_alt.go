package skpsilk

// silk/src/SKP_Silk_biquad_alt.c

func biquad_alt(in []int16, B_Q28 []int32, A_Q28 []int32, S *[]int32, len int32) (out []int16) {
	A0_L_Q28 := (-A_Q28[0]) & 0x3FFF
	A0_U_Q28 := (-A_Q28[0]) >> 14
	A1_L_Q28 := (-A_Q28[1]) & 0x3FFF
	A1_U_Q28 := (-A_Q28[1]) >> 14

	out = make([]int16, len)

	for k := int32(0); k < len; k++ {
		inval := int32(in[k])
		out32_Q14 := SMLAWB((*S)[0], B_Q28[0], inval) << 2

		(*S)[0] = (*S)[1] + RSHIFT_ROUND(SMULWB(out32_Q14, A0_L_Q28), 14)
		(*S)[0] = SMLAWB((*S)[0], out32_Q14, A0_U_Q28)
		(*S)[0] = SMLAWB((*S)[0], B_Q28[1], inval)

		(*S)[1] = RSHIFT_ROUND(SMULWB(out32_Q14, A1_L_Q28), 14)
		(*S)[1] = SMLAWB((*S)[1], out32_Q14, A1_U_Q28)
		(*S)[1] = SMLAWB((*S)[1], B_Q28[2], inval)

		out[k] = SAT16(int16((out32_Q14 + (1 << 14) - 1) >> 14))
	}

	return
}
