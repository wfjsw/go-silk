package skpsilk

// silk/src/SKP_Silk_NLSF2A_stable.c

func NLSF2A_stable(pAR_Q12 []int16, pNLSF []int, LPC_order int) {
	var invGain_Q30 int32
	NLSF2A(pAR_Q12, pNLSF, LPC_order)

	for i := 0; i < MAX_LPC_STABILIZE_ITERATIONS; i++ {
		if LPC_inverse_pred_gain(&invGain_Q30, pAR_Q12, LPC_order) == 1 {
			bwexpander(pAR_Q12, LPC_order, 65536-SMULBB(int32(10+i), int32(i)))
		} else {
			break
		}
	}

	if i == MAX_LPC_STABILIZE_ITERATIONS {
		for i := 0; i < LPC_order; i++ {
			pAR_Q12[i] = 0
		}
	}
}
