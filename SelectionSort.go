package Sort

import (
	"DbxDev/Logger"
)

// A pointer from left to right.
// All element at the left of the pointer are sorted
func SelectionSort(values Sortable) {
	for i := 1; i < values.Size(); i++ {
		for j := i; j > 0; j-- {
			left_j := j - 1
			Logger.Debugf("Step : left_j=%v,j=%v - %v", left_j, j, values.Less(j, left_j))
			if values.Less(j, left_j) {
				values.Swap(left_j, j)
			}
		}
	}
}
