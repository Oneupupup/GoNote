package test6

import (
	"runtime"
	_"fmt"
)

func TestRuntime() {
	if runtime.NumCPU() > 7 {
		runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	}
	runtime.Goexit()
}