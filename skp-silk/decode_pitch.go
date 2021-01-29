package skpsilk

// silk/src/SKP_Silk_decode_pitch.go

func decode_pitch(lagIndex, contourIndex int, pitch_lags []int, Fs_kHz int) {
	min_lag := int(SMULBB(PITCH_EST_MIN_LAG_MS, int32(Fs_kHz)))

	lag := min_lag + lagIndex

	if Fs_kHz == 8 {
		for i := 0; i < PITCH_EST_NB_SUBFR; i++ {
			pitch_lags[i] = lag + CB_lags_stage2[i][contourIndex]
		}
	} else {
		for i := 0; i < PITCH_EST_NB_SUBFR; i++ {
			pitch_lags[i] = lag + CB_lags_stage3[i][contourIndex]
		}
	}
}
