package Sort

import (
	"DbxDev/Logger"
	"math/rand"
	"testing"
)

var skipAll = false

func TestInitShell(t *testing.T) {
	Logger.Init()
	Logger.SetInfo()
	skipAll = true
}
func doOrSkip(t *testing.T) {
	if skipAll {
		t.SkipNow()
	}
}
func TestShellSort(t *testing.T) {
	doOrSkip(t)
	var a IntArray = make([]int, 10)
	h := 3
	for i := 0; i < len(a); i++ {
		a[i] = i % (h - 1)
	}
	Logger.Debugf("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Debugf("After shell-%v sorting \n%v", h, a)
}
func TestBigShellSort(t *testing.T) {
	doOrSkip(t)
	var a IntArray = make([]int, 100)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % 100
	}
	h := 13
	Logger.Debugf("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Debugf("After shell-%v sorting \n%v", h, a)

	h = 7
	Logger.Debugf("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Debugf("After shell-%v sorting \n%v", h, a)

	h = 3
	Logger.Debugf("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Debugf("After shell-%v sorting \n%v", h, a)

	h = 1
	Logger.Debugf("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Debugf("After shell-%v sorting \n%v", h, a)

	if a.IsSorted() {
		Logger.Infof("Array is (shell)sorted")
	} else {
		t.Errorf("shell of heap=1 sort failed")
	}
}
