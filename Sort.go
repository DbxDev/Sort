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

// Merge sort on int array
//
func MergeSort(values *IntArray) {
	lo := 0
	hi := values.Size() - 1
	var aux IntArray = make([]int, values.Size())
	mergeSort(values, &aux, lo, hi)
}

// Recurcive method using an auxiliary array
func mergeSort(values *IntArray, aux *IntArray, lo, hi int) {

	Logger.Debugf("Merge sort with lo=%v , hi=%v", lo, hi)
	mid := lo + (hi-lo)/2
	// if there is at least 3 elements we mergesort again
	if lo < mid && mid < hi {
		mergeSort(values, aux, lo, mid)
		mergeSort(values, aux, mid+1, hi)
	}
	merge(values, aux, lo, mid, hi)
}

// Sort [lo,mid[ and [mid, hi[ parts of an array to an auxiliary
// array
func merge(values *IntArray, aux *IntArray, lo, mid, hi int) {
	// i pointer on [lo,mid[ and j pointer on [mid, hi]
	i, j := lo, mid+1
	k := lo
	Logger.Debugf("{merge}: lo=%v , mid=%v , hi=%v", lo, mid, hi)
	for k <= hi {
		Logger.Debugf("i=%v j=%v k=%v", i, j, k)
		// right part is consummed
		if j > hi {
			Logger.Debugf("{merge}: #1")
			(*aux)[k] = (*values)[i]
			i++
			// left part is consummed
		} else if i > mid {
			Logger.Debugf("{merge}: #2")
			(*aux)[k] = (*values)[j]
			j++
			// left part is smaller than the right part
		} else if values.Less(i, j) {
			Logger.Debugf("{merge}: #3")
			(*aux)[k] = (*values)[i]
			i++
			// right part is smaller than the left part
		} else {
			Logger.Debugf("{merge}: #4")
			(*aux)[k] = (*values)[j]
			j++
		}
		k++
		Logger.Debugf("[%v] extract aux array : %v", k, (*aux)[lo:hi])
	}
	for i := 0; i <= hi-lo; i++ {
		(*values)[lo+i] = (*aux)[lo+i]
	}
}

func QuickSort(values *IntArray) {

}
