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
	Logger.Infof("After quick sorting %v", a)

}
func TestQuickSortBig(t *testing.T) {
	//t.SkipNow()
	Logger.SetInfo()
	size := 5000
	var a IntArray = make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = rand.Intn((size * 10) / 100)
	}
	QuickSort(&a)
	if !a.IsSorted() {
		t.Errorf("Quick sort failed")
	}
	Logger.Infof("Big array of size %v is sorted.", size)

}
