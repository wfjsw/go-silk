package skpsilk

// silk/src/SKP_Silk_ana_filt_bank_1.c

import "math"

/* Coefficients for 2-band filter bank based on first-order allpass filters */
// old
var A_fb1_20 = []int16{5394 << 1}
var A_fb1_21 = []int16{-24290} // {20623 << 1} /* wrap-around to negative number is intentional */

/* Split signal into two decimated bands using first-order allpass filters */
func ana_filt_bank_1(in []int16, S *[]int32, outL *[]int16, outH *[]int16, scratch []int32, N int32) {
	N2 := int(N >> 1)
	for k := 0; k < N2; k++ {
		/* Convert to Q10 */
		in32 := int32(in[2*k]) << 10

		/* All-pass section for even input sample */
		Y := in32 - (*S)[0]
		X := Y + ((Y>>16)*int32(int16((A_fb1_21[0]))) + (((Y & 0x0000FFFF) * int32(int16(A_fb1_21[0]))) >> 16))
		out_1 := (*S)[0] + X
		(*S)[0] = in32 + X

		/* Convert to Q10 */
		in32 = int32(in[2*k+1]) << 10

		/* All-pass section for odd input sample */
		Y = in32 - (*S)[1]
		X = (((Y >> 16) * int32(int16(A_fb1_20[0]))) + (((Y & 0x0000FFFF) * int32(int16(A_fb1_20[0]))) >> 16))
		out_2 := (*S)[1] + X
		(*S)[1] = in32 + X

		/* Add/subtract, convert back to int16 and store to output */
		if int16((((out_2+out_1)>>(11-1))+1)>>1) > math.MaxInt16 {
			(*outL)[k] = math.MaxInt16
		} else if ((((out_2 + out_1) >> (11 - 1)) + 1) >> 1) < math.MinInt16 {
			(*outL)[k] = math.MinInt16
		} else {
			(*outL)[k] = int16((((out_2 + out_1) >> (11 - 1)) + 1) >> 1)
		}

		if int16((((out_2-out_1)>>(11-1))+1)>>1) > math.MaxInt16 {
			(*outH)[k] = math.MaxInt16
		} else if int16((((out_2-out_1)>>(11-1))+1)>>1) < math.MinInt16 {
			(*outH)[k] = math.MinInt16
		} else {
			(*outH)[k] = int16((((out_2 - out_1) >> (11 - 1)) + 1) >> 1)
		}
	}
}
