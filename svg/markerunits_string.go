// Code generated by "stringer -type=MarkerUnits"; DO NOT EDIT.

package svg

import (
	"fmt"
	"strconv"
)

const _MarkerUnits_name = "StrokeWidthUserSpaceOnUseMarkerUnitsN"

var _MarkerUnits_index = [...]uint8{0, 11, 25, 37}

func (i MarkerUnits) String() string {
	if i < 0 || i >= MarkerUnits(len(_MarkerUnits_index)-1) {
		return "MarkerUnits(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MarkerUnits_name[_MarkerUnits_index[i]:_MarkerUnits_index[i+1]]
}

func (i *MarkerUnits) FromString(s string) error {
	for j := 0; j < len(_MarkerUnits_index)-1; j++ {
		if s == _MarkerUnits_name[_MarkerUnits_index[j]:_MarkerUnits_index[j+1]] {
			*i = MarkerUnits(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type MarkerUnits", s)
}
