// Code generated by "stringer -type=Env -output=env_string.go"; DO NOT EDIT.

package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[test-0]
	_ = x[local-1]
	_ = x[dev-2]
	_ = x[prod-3]
}

const _Env_name = "testlocaldevprod"

var _Env_index = [...]uint8{0, 4, 9, 12, 16}

func (i Env) String() string {
	if i < 0 || i >= Env(len(_Env_index)-1) {
		return "Env(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Env_name[_Env_index[i]:_Env_index[i+1]]
}
