// Code generated by "stringer -type=TimeInForce"; DO NOT EDIT.

package robinhood

import "strconv"

const _TimeInForce_name = "GFDGTCIOCOPG"

var _TimeInForce_index = [...]uint8{0, 3, 6, 9, 12}

func (i TimeInForce) String() string {
	if i < 0 || i >= TimeInForce(len(_TimeInForce_index)-1) {
		return "TimeInForce(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TimeInForce_name[_TimeInForce_index[i]:_TimeInForce_index[i+1]]
}
