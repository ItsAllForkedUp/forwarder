// Code generated by "stringer -type=Mode"; DO NOT EDIT.

package forwarder

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Direct-0]
	_ = x[Upstream-1]
	_ = x[PAC-2]
}

const _Mode_name = "DirectUpstreamPAC"

var _Mode_index = [...]uint8{0, 6, 14, 17}

func (i Mode) String() string {
	if i >= Mode(len(_Mode_index)-1) {
		return "Mode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Mode_name[_Mode_index[i]:_Mode_index[i+1]]
}