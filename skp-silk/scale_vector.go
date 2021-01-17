package skpsilk

// silk/src/SKP_Silk_scale_vector.c

// scale_vector32_Q26_lshift_18 Multiply a vector by a constant
func scale_vector32_Q26_lshift_18(data1 []int32, gain_Q26 int32, dataSize int) {
	for i := 0; i < dataSize; i++ {
		data1[i] = int32((int64(data1[i]) * int64(gain_Q26)) >> 8) // OUTPUT: Q18
	}
}
