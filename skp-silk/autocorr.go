package skpsilk

import (
	pmath "github.com/pkg/math"
)

// silk/src/SKP_Silk_autocorr.c

func autocurr(results *[]int32, scale *int, inputData []int16, inputDataSize int, correlationCount int) {
	corrCount := pmath.MinInt(inputDataSize, correlationCount)

	corr64 := inner_prod16_aligned_64(inputData, inputData, inputDataSize)

	/* deal with all-zero input data */
	corr64 += 1

	lz := CLZ64(corr64)

	nRightShifts := 35 - lz
	scale = nRightShifts

	absNRightShifts := nRightShifts
	if absNRightShifts < 0 {
		absNRightShifts = -absNRightShifts
	}

	(*results)[0] = int32(corr64) << absNRightShifts

	/* compute remaining correlations based on int32 inner product */
	for i := 1; i < corrCount; i++ {
		(*results)[i] = inner_prod_aligned(inputData, inputData+i, inputData-i) << absNRightShifts
	}
}
