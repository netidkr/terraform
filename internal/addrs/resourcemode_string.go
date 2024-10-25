// Code generated by "stringer -type ResourceMode"; DO NOT EDIT.

package addrs

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidResourceMode-0]
	_ = x[ManagedResourceMode-77]
	_ = x[DataResourceMode-68]
	_ = x[EphemeralResourceMode-69]
}

const (
	_ResourceMode_name_0 = "InvalidResourceMode"
	_ResourceMode_name_1 = "DataResourceModeEphemeralResourceMode"
	_ResourceMode_name_2 = "ManagedResourceMode"
)

var (
	_ResourceMode_index_1 = [...]uint8{0, 16, 37}
)

func (i ResourceMode) String() string {
	switch {
	case i == 0:
		return _ResourceMode_name_0
	case 68 <= i && i <= 69:
		i -= 68
		return _ResourceMode_name_1[_ResourceMode_index_1[i]:_ResourceMode_index_1[i+1]]
	case i == 77:
		return _ResourceMode_name_2
	default:
		return "ResourceMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}