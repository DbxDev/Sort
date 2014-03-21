package Sort

import (
	"DbxDev/Logger"
	"math/rand"
	"testing"
)

func TestInitQuickSort(t *testing.T) {
	Logger.Init()
	Logger.SetDebug()
}

func TestQuickSortSmall(t *testing.T) {
	size := 10
	var a IntArray = make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(size / 3)
	}
	Logger.Infof("Before insert sorting %v", a)
	QuickSort(&a)
	if !a.IsSorted() {
		t.Errorf("Quick sort failed")
	}
	Logger.Infof("After insertion sorting %v", a)

}
func TestQuickSortBig(t *testing.T) {
	t.SkipNow()
	Logger.SetInfo()
	size := 50
	var a IntArray = make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn((size * 10) / 100)
	}
	Logger.Infof("Before insert sorting %v", a)
	QuickSort(&a)
	if !a.IsSorted() {
		t.Errorf("Quick sort failed")
	}
	Logger.Infof("After insertion sorting %v", a)

}
