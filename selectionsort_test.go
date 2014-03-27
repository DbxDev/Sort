package Sort

import (
	"DbxDev/Logger"
	"math/rand"
	"testing"
)

func TestInitSelect(t *testing.T) {
	Logger.Init()
	Logger.SetInfo()
}
func TestIsSorted(t *testing.T) {
	var unsorted IntArray = []int{0, 1, 1, 2, 3, 5, 4, 6, 7}
	if unsorted.IsSorted() {
		t.Errorf("Unsorted array : expected IsSorted to return false")
	}
	var sorted IntArray = []int{0, 1, 1, 2, 3, 3, 4, 6, 7}
	if !sorted.IsSorted() {
		t.Errorf("Sorted array : expected IsSorted to return true")
	}
}
func TestSmallSelectionSort(t *testing.T) {
	t.SkipNow()
	var a IntArray = make([]int, 10)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % (i + 1)
	}
	Logger.Infof("Before Selection sorting %v", a)
	SelectionSort(a)
	if !a.IsSorted() {
		t.Errorf("Selection sort failed")
	}
	Logger.Infof("After sorting %v", a)
}
func TestBigSelectionSort(t *testing.T) {
	t.SkipNow()
	bigSize := 1000
	var a IntArray = make([]int, bigSize)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % (100)
	}
	if a.IsSorted() {
		Logger.Infof("Array of size %v is sorted", bigSize)
	} else {
		Logger.Infof("Array of size %v is unsorted", bigSize)
	}

	SelectionSort(a)
	if !a.IsSorted() {
		t.Errorf("Selection sort failed")
	}
	Logger.Infof("Array of size %v is (selection)sorted", bigSize)
}
