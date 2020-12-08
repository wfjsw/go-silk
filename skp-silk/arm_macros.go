package skpsilk

// silk/src/SKP_Silk_macros.h

func SMULWB(a32, b32 int32) int32 {
	return (a32>>16)*(b32&0xFFFF) + (((a32 & 0xFFFF) * (b32 & 0xFFFF)) >> 16)
}

func SMLAWB(a32, b32, c32 int32) int32 {
	return a32 + (b32>>16)*(c32&0xFFFF) + (((b32 & 0xFFFF) * (c32 & 0xFFFF)) >> 16)
}

func SMULWT(a32, b32 int32) int32 {
	return (a32>>16)*(b32>>16) + (((a32 & 0xFFFF) * (b32 >> 16)) >> 16)
}

func SMLAWT(a32, b32, c32 int32) int32 {
	return a32 + (b32>>16)*(c32>>16) + (((b32 & 0xFFFF) * (c32 >> 16)) >> 16)
}

func SMULBB(a32, b32 int32) int32 {
	return (a32 & 0xFFFF) * (b32 & 0xFFFF)
}

func SMLABB(a32, b32, c32 int32) int32 {
	return a32 + (b32&0xFFFF)*(c32&0xFFFF)
}

func SMLABT(a32, b32, c32 int32) int32 {
	return a32 + (b32&0xFFFF)*(c32>>16)
}

func SMLAL(a64 int64, b32, c32 int32) int64 {
	return (a64) + int64(b32)*int64(c32)
}

func SMULWW(a32, b32 int32) int32 {
	return SMULWB(a32, b32) + a32*RSHIFT_ROUND(b32, 16)
}

func SMLAWW(a32, b32, c32 int32) int32 {
	return SMLAWB(a32, b32, c32) + b32*RSHIFT_ROUND(c32, 16)
}

func SMMUL(a32, b32 int32) int32 {
	return int32((int64(a32) * int64(b32)) >> 32)
}

func CLZ16(in16 int16) int32 {
	out32 := int32(0)
	if in16 == 0 {
		return 16
	}
	/* test nibbles */
	if in16&-256 != 0 {
		if in16&-4096 != 0 {
			in16 >>= 12
		} else {
			out32 += 4
			in16 >>= 8
		}
	} else {
		if in16&-16 != 0 {
			out32 += 8
			in16 >>= 4
		} else {
			out32 += 12
		}
	}
	/* test bits and return */
	if in16&0xC != 0 {
		if in16&0x8 != 0 {
			return out32 + 0
		} else {
			return out32 + 1
		}
	} else {
		if in16&0xE != 0 {
			return out32 + 2
		} else {
			return out32 + 3
		}
	}
}

func CLZ32(in32 int32) int32 {
	if in32&-65536 != 0 {
		return CLZ16(int16(in32 >> 16))
	} else {
		return CLZ16(int16(in32)) + 16
	}
}
