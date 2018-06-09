// Code generated by "stringer -type=Visibilities"; DO NOT EDIT.

package window

import (
	"fmt"
	"strconv"
)

const _Visibilities_name = "ClosedNotVisibleVisibleIconifiedVisibilitiesN"

var _Visibilities_index = [...]uint8{0, 6, 16, 23, 32, 45}

func (i Visibilities) String() string {
	if i < 0 || i >= Visibilities(len(_Visibilities_index)-1) {
		return "Visibilities(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Visibilities_name[_Visibilities_index[i]:_Visibilities_index[i+1]]
}

func (i *Visibilities) FromString(s string) error {
	for j := 0; j < len(_Visibilities_index)-1; j++ {
		if s == _Visibilities_name[_Visibilities_index[j]:_Visibilities_index[j+1]] {
			*i = Visibilities(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type Visibilities", s)
}