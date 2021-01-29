package skpsilk

// silk/interface/SKP_Silk_control.h

// EncControlStruct Structure for controlling encoder operation
type EncControlStruct struct {
	// I:   Input signal sampling rate in Hertz; 8000/12000/16000/24000
	API_sampleRate int32
	// I:   Maximum internal sampling rate in Hertz; 8000/12000/16000/24000
	maxInternalSampleRate int32
	// I:   Number of samples per packet; must be equivalent of 20, 40, 60, 80 or 100 ms
	packetSize           int
	bitRate              int32
	packetLossPercentage int
	complexity           int
	useInBandFEC         int
	useDTX               int
}
