package skpsilk

// silk/src/SKP_Silk_define.h

const (
	MAX_FRAMES_PER_PACKET = 5

	/* Limits on bitrate */
	MIN_TARGET_RATE_BPS = 5000
	MAX_TARGET_RATE_BPS = 100000

	/* Transition bitrates between modes */
	SWB2WB_BITRATE_BPS = 25000
	WB2SWB_BITRATE_BPS = 30000
	WB2MB_BITRATE_BPS  = 14000
	MB2WB_BITRATE_BPS  = 18000
	MB2NB_BITRATE_BPS  = 10000
	NB2MB_BITRATE_BPS  = 14000

	/* Integration/hysteresis threshold for lowering internal sample frequency */
	/* 30000000 -> 6 sec if bitrate is 5000 bps below limit; 3 sec if bitrate is 10000 bps below limit */
	ACCUM_BITS_DIFF_THRESHOLD = 30000000
	TARGET_RATE_TAB_SZ        = 8

	/* DTX settings                                 */
	NO_SPEECH_FRAMES_BEFORE_DTX = 5  /* eq 100 ms */
	MAX_CONSECUTIVE_DTX         = 20 /* eq 400 ms */

	USE_LBRR = 1

	/* Amount of concecutive no FEC packets before telling JB */
	NO_LBRR_THRES = 10

	/* Maximum delay between real packet and LBRR packet */
	MAX_LBRR_DELAY = 2
	LBRR_IDX_MASK  = 1

	INBAND_FEC_MIN_RATE_BPS = 18000 /* Dont use inband FEC below this total target rate  */
	LBRR_LOSS_THRES         = 1     /* Start adding LBRR at this loss rate                  */

	/* LBRR usage defines */
	SKP_SILK_NO_LBRR           = 0 /* No LBRR information for this packet                  */
	SKP_SILK_ADD_LBRR_TO_PLUS1 = 1 /* Add LBRR for this packet to packet n + 1             */
	SKP_SILK_ADD_LBRR_TO_PLUS2 = 2 /* Add LBRR for this packet to packet n + 2             */

	/* Frame termination indicator defines */
	SKP_SILK_LAST_FRAME  = 0 /* Last frames in packet                                */
	SKP_SILK_MORE_FRAMES = 1 /* More frames to follow this one                       */
	SKP_SILK_LBRR_VER1   = 2 /* LBRR information from packet n - 1                   */
	SKP_SILK_LBRR_VER2   = 3 /* LBRR information from packet n - 2                   */
	SKP_SILK_EXT_LAYER   = 4 /* Extension layers added                               */

	/* Number of Second order Sections for SWB detection HP filter */
	NB_SOS                           = 3
	HP_8_KHZ_THRES                   = 10       /* average energy per sample, above 8 kHz       */
	CONCEC_SWB_SMPLS_THRES           = 480 * 15 /* 300 ms                                       */
	WB_DETECT_ACTIVE_SPEECH_MS_THRES = 15000    /* ms of active speech needed for WB detection  */

	/* Low complexity setting */
	LOW_COMPLEXITY_ONLY = 0

	/* Activate bandwidth transition filtering for mode switching */
	SWITCH_TRANSITION_FILTERING = 1

	/* Decoder Parameters */
	DEC_HP_ORDER = 2

	/* Maximum sampling frequency, should be 16 for some embedded platforms */
	MAX_FS_KHZ     = 24
	MAX_API_FS_KHZ = 48

	/* Signal Types used by silk */
	SIG_TYPE_VOICED   = 0
	SIG_TYPE_UNVOICED = 1

	/* VAD Types used by silk */
	NO_VOICE_ACTIVITY = 0
	VOICE_ACTIVITY    = 1

	/* Number of samples per frame */
	FRAME_LENGTH_MS  = 20
	MAX_FRAME_LENGTH = FRAME_LENGTH_MS * MAX_FS_KHZ

	/* Milliseconds of lookahead for pitch analysis */
	LA_PITCH_MS  = 2
	LA_PITCH_MAX = LA_PITCH_MS * MAX_FS_KHZ

	/* Length of LPC window used in find pitch */
	FIND_PITCH_LPC_WIN_MS  = 20 + (LA_PITCH_MS << 1)
	FIND_PITCH_LPC_WIN_MAX = FIND_PITCH_LPC_WIN_MS * MAX_FS_KHZ

	/* Order of LPC used in find pitch */
	MAX_FIND_PITCH_LPC_ORDER = 16

	/* Pitch estimator */
	PITCH_EST_MIN_COMPLEX = 0
	PITCH_EST_MID_COMPLEX = 1
	PITCH_EST_MAX_COMPLEX = 2

	PITCH_EST_COMPLEXITY_HC_MODE = PITCH_EST_MAX_COMPLEX
	PITCH_EST_COMPLEXITY_MC_MODE = PITCH_EST_MID_COMPLEX
	PITCH_EST_COMPLEXITY_LC_MODE = PITCH_EST_MIN_COMPLEX

	/* Milliseconds of lookahead for noise shape analysis */
	LA_SHAPE_MS  = 5
	LA_SHAPE_MAX = (LA_SHAPE_MS * MAX_FS_KHZ)

	/* Max length of LPC window used in noise shape analysis */
	SHAPE_LPC_WIN_MAX = (15 * MAX_FS_KHZ)

	/* Max number of bytes in payload output buffer (may contain multiple frames) */
	MAX_ARITHM_BYTES = 1024

	RANGE_CODER_WRITE_BEYOND_BUFFER   = -1
	RANGE_CODER_CDF_OUT_OF_RANGE      = -2
	RANGE_CODER_NORMALIZATION_FAILED  = -3
	RANGE_CODER_ZERO_INTERVAL_WIDTH   = -4
	RANGE_CODER_DECODER_CHECK_FAILED  = -5
	RANGE_CODER_READ_BEYOND_BUFFER    = -6
	RANGE_CODER_ILLEGAL_SAMPLING_RATE = -7
	RANGE_CODER_DEC_PAYLOAD_TOO_LONG  = -8

	/* dB level of lowest gain quantization level */
	MIN_QGAIN_DB = 6
	/* dB level of highest gain quantization level */
	MAX_QGAIN_DB = 86
	/* Number of gain quantization levels */
	N_LEVELS_QGAIN = 64
	/* Max increase in gain quantization index */
	MAX_DELTA_GAIN_QUANT = 40
	/* Max decrease in gain quantization index */
	MIN_DELTA_GAIN_QUANT = -4

	/* Quantization offsets (multiples of 4) */
	OFFSET_VL_Q10  = 32
	OFFSET_VH_Q10  = 100
	OFFSET_UVL_Q10 = 100
	OFFSET_UVH_Q10 = 256

	/* Maximum numbers of iterations used to stabilize a LPC vector */
	MAX_LPC_STABILIZE_ITERATIONS = 20

	MAX_LPC_ORDER = 16
	MIN_LPC_ORDER = 10

	/* Find Pred Coef defines */
	LTP_ORDER = 5

	/* LTP quantization settings */
	NB_LTP_CBKS = 3

	/* Number of subframes */
	NB_SUBFR = 4

	/* Flag to use harmonic noise shaping */
	USE_HARM_SHAPING = 1

	/* Max LPC order of noise shaping filters */
	MAX_SHAPE_LPC_ORDER = 16

	HARM_SHAPE_FIR_TAPS = 3

	/* Maximum number of delayed decision states */
	MAX_DEL_DEC_STATES = 4

	LTP_BUF_LENGTH = 512
	LTP_MASK       = (LTP_BUF_LENGTH - 1)

	DECISION_DELAY      = 32
	DECISION_DELAY_MASK = (DECISION_DELAY - 1)

	/* number of subframes for excitation entropy coding */
	SHELL_CODEC_FRAME_LENGTH = 16
	MAX_NB_SHELL_BLOCKS      = (MAX_FRAME_LENGTH / SHELL_CODEC_FRAME_LENGTH)

	/* number of rate levels, for entropy coding of excitation */
	N_RATE_LEVELS = 10

	/* maximum sum of pulses per shell coding frame */
	MAX_PULSES = 18

	MAX_MATRIX_SIZE = MAX_LPC_ORDER /* Max of LPC Order and LTP order */

	// #if( MAX_LPC_ORDER > DECISION_DELAY )
	// # define NSQ_LPC_BUF_LENGTH                     MAX_LPC_ORDER
	// #else
	NSQ_LPC_BUF_LENGTH = DECISION_DELAY
	// #endif

	/***********************/
	/* High pass filtering */
	/***********************/
	HIGH_PASS_INPUT = 1

	/***************************/
	/* Voice activity detector */
	/***************************/
	VAD_N_BANDS = 4

	VAD_INTERNAL_SUBFRAMES_LOG2 = 2
	VAD_INTERNAL_SUBFRAMES      = (1 << VAD_INTERNAL_SUBFRAMES_LOG2)

	VAD_NOISE_LEVEL_SMOOTH_COEF_Q16 = 1024 /* Must be < 4096                                   */
	VAD_NOISE_LEVELS_BIAS           = 50

	/* Sigmoid settings */
	VAD_NEGATIVE_OFFSET_Q5 = 128 /* sigmoid is 0 at -128                             */
	VAD_SNR_FACTOR_Q16     = 45000

	/* smoothing for SNR measurement */
	VAD_SNR_SMOOTH_COEF_Q18 = 4096

	/******************/
	/* NLSF quantizer */
	/******************/
	NLSF_MSVQ_MAX_CB_STAGES                   = 10  /* Update manually when changing codebooks      */
	NLSF_MSVQ_MAX_VECTORS_IN_STAGE            = 128 /* Update manually when changing codebooks      */
	NLSF_MSVQ_MAX_VECTORS_IN_STAGE_TWO_TO_END = 16  /* Update manually when changing codebooks      */

	NLSF_MSVQ_FLUCTUATION_REDUCTION = 1
	MAX_NLSF_MSVQ_SURVIVORS         = 16
	MAX_NLSF_MSVQ_SURVIVORS_LC_MODE = 2
	MAX_NLSF_MSVQ_SURVIVORS_MC_MODE = 4

	/* Based on above defines, calculate how much memory is necessary to allocate */
	// #if( NLSF_MSVQ_MAX_VECTORS_IN_STAGE > ( MAX_NLSF_MSVQ_SURVIVORS_LC_MODE * NLSF_MSVQ_MAX_VECTORS_IN_STAGE_TWO_TO_END ) )
	NLSF_MSVQ_TREE_SEARCH_MAX_VECTORS_EVALUATED_LC_MODE = NLSF_MSVQ_MAX_VECTORS_IN_STAGE
	// #else
	// #   define NLSF_MSVQ_TREE_SEARCH_MAX_VECTORS_EVALUATED_LC_MODE  MAX_NLSF_MSVQ_SURVIVORS_LC_MODE * NLSF_MSVQ_MAX_VECTORS_IN_STAGE_TWO_TO_END
	// #endif

	// #if( NLSF_MSVQ_MAX_VECTORS_IN_STAGE > ( MAX_NLSF_MSVQ_SURVIVORS * NLSF_MSVQ_MAX_VECTORS_IN_STAGE_TWO_TO_END ) )
	// #   define NLSF_MSVQ_TREE_SEARCH_MAX_VECTORS_EVALUATED  NLSF_MSVQ_MAX_VECTORS_IN_STAGE
	// #else
	NLSF_MSVQ_TREE_SEARCH_MAX_VECTORS_EVALUATED = MAX_NLSF_MSVQ_SURVIVORS * NLSF_MSVQ_MAX_VECTORS_IN_STAGE_TWO_TO_END
	// #endif

	NLSF_MSVQ_SURV_MAX_REL_RD = 0.1

	// #if SWITCH_TRANSITION_FILTERING
	TRANSITION_TIME_UP_MS     = 5120 // 5120 = 64 * FRAME_LENGTH_MS * ( TRANSITION_INT_NUM - 1 ) = 64*(20*4)
	TRANSITION_TIME_DOWN_MS   = 2560 // 2560 = 32 * FRAME_LENGTH_MS * ( TRANSITION_INT_NUM - 1 ) = 32*(20*4)
	TRANSITION_NB             = 3    /* Hardcoded in tables */
	TRANSITION_NA             = 2    /* Hardcoded in tables */
	TRANSITION_INT_NUM        = 5    /* Hardcoded in tables */
	TRANSITION_FRAMES_UP      = (TRANSITION_TIME_UP_MS / FRAME_LENGTH_MS)
	TRANSITION_FRAMES_DOWN    = (TRANSITION_TIME_DOWN_MS / FRAME_LENGTH_MS)
	TRANSITION_INT_STEPS_UP   = (TRANSITION_FRAMES_UP / (TRANSITION_INT_NUM - 1))
	TRANSITION_INT_STEPS_DOWN = (TRANSITION_FRAMES_DOWN / (TRANSITION_INT_NUM - 1))
	// #endif

	/* BWE factors to apply after packet loss */
	BWE_AFTER_LOSS_Q16 = 63570

	/* Defines for CN generation */
	CNG_BUF_MASK_MAX  = 255
	CNG_GAIN_SMTH_Q16 = 4634
	CNG_NLSF_SMTH_Q16 = 16348
)

// /* Row based */
// #define matrix_ptr(Matrix_base_adr, row, column, N)         *(Matrix_base_adr + ((row)*(N)+(column)))
// #define matrix_adr(Matrix_base_adr, row, column, N)          (Matrix_base_adr + ((row)*(N)+(column)))

// /* Column based */
// #ifndef matrix_c_ptr
// #   define matrix_c_ptr(Matrix_base_adr, row, column, M)    *(Matrix_base_adr + ((row)+(M)*(column)))
// #endif
// #define matrix_c_adr(Matrix_base_adr, row, column, M)        (Matrix_base_adr + ((row)+(M)*(column)))
