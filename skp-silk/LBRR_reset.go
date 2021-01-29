package skpsilk

// silk/src/SKP_Silk_LBRR_reset.c

// LBRR_reset Resets LBRR buffer, used if packet size changes
func LBRR_reset(psEncC *encoder_state) {
	for i := 0; i < MAX_LBRR_DELAY; i++ {
		psEncC.LBRR_buffer[i].usage = NO_LBRR
	}
}
