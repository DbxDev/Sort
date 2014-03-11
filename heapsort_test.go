package Sort

import (
	"DbxDev/Logger"
	"math/rand"
	"testing"
)

func TestInitShell(t *testing.T) {
	Logger.Init()
	Logger.SetInfo()
}

func TestShellSort(t *testing.T) {
	var a IntArray = make([]int, 10)
	h := 3
	for i := 0; i < len(a); i++ {
		a[i] = i % (h - 1)
	}
	Logger.Infof("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Infof("After shell-%v sorting \n%v", h, a)
}
func TestBigShellSort(t *testing.T) {
	var a IntArray = make([]int, 100)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % 100
	}
	h := 13
	Logger.Infof("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Infof("After shell-%v sorting \n%v", h, a)

	h = 7
	Logger.Infof("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Infof("After shell-%v sorting \n%v", h, a)

	h = 3
	Logger.Infof("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Infof("After shell-%v sorting \n%v", h, a)

	h = 1
	Logger.Infof("Before shell-%v sorting \n%v", h, a)
	ShellSort(a, h)
	Logger.Infof("After shell-%v sorting \n%v", h, a)

	if a.IsSorted() {
		Logger.Infof("Array is (shell)sorted")
	} else {
		t.Errorf("shell of heap=1 sort failed")
	}
}
