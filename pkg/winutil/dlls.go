package winutil

import "golang.org/x/sys/windows"

// all dlls used by winutil
var (
	dllKernel32 = windows.NewLazyDLL("kernel32.dll")
)

// all dll exports used by winutil
var (
	winapiReadProcessMemory = dllKernel32.NewProc("ReadProcessMemory")
)
