// Code generated by "stringer -type=ErrorCode"; DO NOT EDIT.

package irtt

import "fmt"

const _ErrorCode_name = "InvalidWinAvgWindowInvalidExpAvgAlphaDSCPErrorDFErrorTTLErrorShortWriteExpectedReplyFlagUnexpectedReplyFlagShortReplyStampAtMismatchClockMismatchUnexpectedSequenceNumberInvalidDFStringFieldsLengthTooLargeFieldsCapacityTooLargeInvalidStampAtStringInvalidStampAtIntInvalidAllowStampStringInvalidClockStringInvalidClockIntInvalidSleepFactorIntervalNotPermittedInvalidWaitStringInvalidWaitFactorInvalidWaitDurationNoSuchAveragerNoSuchFillerNoSuchTimerNoSuchWaiterBadMagicNoHMACBadHMACUnexpectedHMACNonexclusiveMidpointTStampInconsistentClocksListenerPanicDFNotSupportedIntervalNonPositiveDurationNonPositiveOpenCloseBothSetConnTokenZeroConnClosedInvalidFlagBitsSetServerClosedShortParamBufferParamOverflowUnknownParamNoSuitableAddressFoundOpenTimeoutInvalidServerRestrictionInvalidParamValueParamsChangedInvalidReceivedStatsIntInvalidReceivedStatsString"

var _ErrorCode_index = [...]uint16{0, 19, 37, 46, 53, 61, 71, 88, 107, 117, 132, 145, 169, 184, 204, 226, 246, 263, 286, 304, 319, 337, 357, 374, 391, 410, 424, 436, 447, 459, 467, 473, 480, 494, 520, 538, 551, 565, 584, 603, 619, 632, 642, 660, 672, 688, 701, 713, 735, 746, 770, 787, 800, 823, 849}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return fmt.Sprintf("ErrorCode(%d)", i)
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
