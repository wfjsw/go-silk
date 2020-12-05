package skpsilk

const (
	MAX_ORDER_LPC          = 16
	MAX_CORRELATION_LENGTH = 640
	LSF_COS_TAB_SZ_FIX     = 128
)

func min_32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
