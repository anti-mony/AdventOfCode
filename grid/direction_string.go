// Code generated by "stringer -type=Direction"; DO NOT EDIT.

package grid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DirectionNorth-1]
	_ = x[DirectionEast-2]
	_ = x[DirectionWest-3]
	_ = x[DirectionSouth-4]
}

const _Direction_name = "DirectionNorthDirectionEastDirectionWestDirectionSouth"

var _Direction_index = [...]uint8{0, 14, 27, 40, 54}

func (i Direction) String() string {
	i -= 1
	if i < 0 || i >= Direction(len(_Direction_index)-1) {
		return "Direction(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Direction_name[_Direction_index[i]:_Direction_index[i+1]]
}