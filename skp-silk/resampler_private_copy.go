package skpsilk

// silk/src/SKP_Silk_resampler_private_copy.go

func resampler_private_copy(_ interface{}, out []int16, in []int16, inLen int32) {
	copy(out, in[:inLen])
}
