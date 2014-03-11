package Sort

import (
	"DbxDev/Logger"
)

// Define a pointer moving from left to right.
// Find the minimum of the right of the pointer (included)
// All the elements at the left of the pointer are sorted
func SelectionSort(values Sortable) {
	for i := 0; i < values.Size(); i++ {
		min := i
		for j := i; j < values.Size(); j++ {
			if values.Less(j, min) {
				min = j
			}
		}
		values.Swap(i, min)
	}
}

/*	Define a pointer moving from left to right
	Slide the pointed value to the left as long as it is smaller
*/
func InsertionSort(values Sortable) {
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

/*	h-sort an array with insertion sort applied
	to every h-th element
*/
func ShellSort(values Sortable, h int) {
	Logger.Infof("Shell sort of heap %v", h)
	for i := 0; i < values.Size(); i++ {
		step := i
		for step >= h {
			if values.Less(step, step-h) {
				values.Swap(step, step-h)
			}
			step -= h
		}
	}
}
