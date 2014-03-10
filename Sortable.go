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

func (a IntArray) Less(i, j int) bool {
	Logger.Debugf("Is a[%v]=%v <= a[%v]=%v", i, a[i], j, a[j])
	return a[i] <= a[j]
}
func (a IntArray) Swap(i, j int) {
	Logger.Debugf("Swap avant : [%v]%v <-> [%v]%v", i, a[i], j, a[j])
	swap := a[i]
	a[i] = a[j]
	a[j] = swap
	Logger.Debugf("Swap après : [%v]%v <-> [%v]%v", i, a[i], j, a[j])
}
func (a IntArray) Size() int {
	return len(a)
}
func (a IntArray) IsSorted() bool {
	for i := 0; i < a.Size()-1; i++ {
		if !a.Less(i, i+1) {
			return false
		}
	}
	return true
}
