// Code generated by "stringer -type=FocusChanges"; DO NOT EDIT.

package gi

import (
	"fmt"
	"strconv"
)

const _FocusChanges_name = "FocusLostFocusGotFocusInactiveFocusActiveFocusChangesN"

var _FocusChanges_index = [...]uint8{0, 9, 17, 30, 41, 54}

func (i FocusChanges) String() string {
	if i < 0 || i >= FocusChanges(len(_FocusChanges_index)-1) {
		return "FocusChanges(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FocusChanges_name[_FocusChanges_index[i]:_FocusChanges_index[i+1]]
}

func (i *FocusChanges) FromString(s string) error {
	for j := 0; j < len(_FocusChanges_index)-1; j++ {
		if s == _FocusChanges_name[_FocusChanges_index[j]:_FocusChanges_index[j+1]] {
			*i = FocusChanges(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type FocusChanges", s)
}
