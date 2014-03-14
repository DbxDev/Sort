package Sort

import (
	"DbxDev/Logger"
	"math/rand"
	"testing"
)

func TestInitMergeSort(t *testing.T) {
	Logger.Init()
	Logger.SetDebug()
	skipAll = false
}

func TestSort(t *testing.T) {
	t.SkipNow()
	doOrSkip(t)
	var a IntArray = make([]int, 20)
	var aux IntArray = make([]int, 20)
	var sub1 IntArray = make([]int, 10)
	var sub2 IntArray = make([]int, 10)
	for i := 0; i < len(sub1); i++ {
		sub1[i] = rand.Int() % 5
		sub2[i] = rand.Int() % 5
	}
	InsertionSort(sub1)
	InsertionSort(sub2)
	Logger.Debugf("Left : %v", sub1)
	Logger.Debugf("Righ : %v", sub2)
	for i := 0; i < 20; i++ {
		if i < 10 {
			a[i] = sub1[i]
		} else {
			a[i] = sub2[i-10]
		}
	}

	lo, hi := 0, len(a)
	mid := lo + (hi-lo)/2
	Logger.Debugf("Array to merge [%v,%v,%v] :\n%v", lo, mid, hi, a)
	merge(&a, &aux, lo, mid, hi)
	Logger.Debugf("Merged array \n%v", aux)
}
func TestSortSmall(t *testing.T) {
	t.SkipNow()
	doOrSkip(t)
	var a IntArray = make([]int, 2)
	var aux IntArray = make([]int, 2)
	a[0], a[1] = 2, 1
	lo, hi := 0, len(a)-1
	mid := lo + (hi-lo)/2
	Logger.Debugf("Array to sort [%v,%v,%v] :\n%v", lo, mid, hi, a)
	merge(&a, &aux, lo, mid, hi)
	Logger.Debugf("Merged array \n%v", a)
}

func TestMergeSort(t *testing.T) {
	//t.SkipNow()
	doOrSkip(t)
	var a IntArray = make([]int, 20)
	for i := 0; i < len(a); i++ {
		//a[i] = rand.Int() % 5
		a[i] = len(a) - i
	}
	Logger.Debugf("Array to Mergesort :\n%v", a)
	MergeSort(&a)
	Logger.Debugf("Merged array \n%v", a)
}
