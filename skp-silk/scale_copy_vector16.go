package skpsilk

// silk/src/SKP_Silk_scale_copy_vector16.c

func scale_copy_vector16(data_out []int16, data_in []int16, gain_Q16 int32, dataSize int) {
	for i := 0; i < dataSize; i++ {
		data_out[i] = int16(SMULWB(int32(gain_Q16), int32(data_in[i])))
	}
}
