package skpsilk

// silk/src/SKP_Silk_regularize_correlations_FIX.c

func regularize_correlations_FIX(XX []int32, xx []int32, noise int32, D int) {
	for i := 0; i < D; i++ {
		XX[D*i+i] += noise
	}
	xx[0] += noise
}
