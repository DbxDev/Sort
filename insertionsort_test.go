package Sort

import (
	"DbxDev/Logger"
	"math/rand"
	"testing"
)

func TestInitInsert(t *testing.T) {
	Logger.Init()
	Logger.SetInfo()
}

func TestSmallInsertionSort(t *testing.T) {
	var a IntArray = make([]int, 10)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % (i + 1)
	}
	Logger.Infof("Before insert sorting %v", a)
	InsertionSort(a)
	if !a.IsSorted() {
		t.Errorf("Insertion sort failed")
	}
	Logger.Infof("After insertion sorting %v", a)
}
func TestBigInsertionSort(t *testing.T) {
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

	InsertionSort(a)
	if !a.IsSorted() {
		t.Errorf("Insertion sort failed")
	}
	Logger.Infof("Array of size %v is (insertion)sorted", bigSize)
}
