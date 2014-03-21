package Sort

import (
	"DbxDev/List"
	"DbxDev/Logger"
	"fmt"
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

// Quicksort a non stable fast sorting method
// It uses a pivot value and parts the array :
// - upper values on the right
// - lower values on the left
func QuickSort(values *IntArray) {
	List.ShuffleInt(*values)
	Logger.Debugf("Shuffled array %v", *values)
	doPivot(values, 0, values.Size()-1)
}

func doPivot(values *IntArray, lo, hi int) {
	if hi < lo {
		panic(fmt.Errorf("hi=%v[%v]<lo=%v[%v]for array %v", hi, (*values)[hi], lo, (*values)[lo], values))
	}
	Logger.Infof("Pivot [%v , %v] - %v", lo, hi, (*values)[lo:hi+1])
	if lo == hi {
		Logger.Debugf("returning...")
		return
	}
	if hi-lo == 1 {
		Logger.Debugf("comparing %v and %v and swap if necessary", lo, hi)
		if values.Less(hi, lo) {
			values.Swap(lo, hi)
		}
		return
	}
	pivot := lo
	i := lo
	j := hi + 1 // so the first j-- make j actually start a 'hi'
	// iterate until pointers cross
	for i < j {
		i++
		j--
		Logger.Debugf("[pivot{%v}=%v] New loop i=%v,j=%v", pivot, (*values)[pivot], i, j)
		for values.Less(i, pivot) && i < j {
			i++
			Logger.Debugf("Left Pointer moved to %v", i)
		}
		for values.Less(pivot, j) && j > i {
			j--
			Logger.Debugf("Right Pointer moved to %v", j)
		}
		Logger.Debugf("Swapping %v - %v", i, j)
		values.Swap(i, j)

		Logger.Debugf("Current array state %v", *values)
	}
	Logger.Debugf("Swapping pivot from %v to %v", pivot, j)
	values.Swap(pivot, j)
	Logger.Debugf("Final state [%v,%v] %v", lo, hi, *values)
	Logger.Infof("Next steps : %v-%v and %v-%v", lo, j-1, j+1, hi)
	doPivot(values, lo, j-1)
	doPivot(values, j+1, hi)
}
