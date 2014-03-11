package Sort

import (
	"DbxDev/Logger"
)

type Sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
	IsSorted() bool
	Size() int
}

type IntArray []int

// Less strictly to ensure stability
func (a IntArray) Less(i, j int) bool {
	return a[i] < a[j]
}
func (a IntArray) Swap(i, j int) {
	if i == j {
		return
	}
	swap := a[i]
	a[i] = a[j]
	a[j] = swap
}
func (a IntArray) Size() int {
	return len(a)
}
func (a IntArray) IsSorted() bool {
	for i := 0; i < a.Size()-1; i++ {
		if a[i] != a[i+1] && !a.Less(i, i+1) {
			Logger.Debugf("Unsorted array")
			return false
		}
	}
	Logger.Debugf("Sorted array")
	return true
}
