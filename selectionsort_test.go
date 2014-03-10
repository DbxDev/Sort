package Sort

import (
	"DbxDev/Logger"
	"testing"
)

func TestIsSorted(t *testing.T) {
	Logger.Init()
	var unsorted IntArray = []int{0, 1, 1, 2, 3, 5, 4, 6, 7}
	if !unsorted.IsSorted() {
		t.Errorf("Unsorted array : expected IsSorted to return false")
	} else {
		Logger.Infof("Unsorted array : %v", unsorted)
	}
	var sorted IntArray = []int{0, 1, 1, 2, 3, 3, 4, 6, 7}
	if sorted.IsSorted() {
		t.Errorf("Sorted array : expected IsSorted to return true")
	} else {
		Logger.Infof("Sorted array : %v", sorted)
	}
}
func TestSelectionSort(t *testing.T) {
	Logger.Init()
	var a IntArray = make([]int, 10)
	for i := 0; i < len(a); i++ {
		a[i] = len(a) - i
	}
	Logger.Debugf("%v", a)
	SelectionSort(a)
	Logger.Debugf("%v", a)
}
