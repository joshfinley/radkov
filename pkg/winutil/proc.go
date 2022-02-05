package winutil

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

type WinProc struct {
	Pid           int            // process id
	LoadedModules []WinMod       // loaded dlls
	Handle        windows.Handle // handle to process
	Path          string         // path of the exe image
}

func NewWinProc(name string) (*WinProc, error) {
	_, err := FindPidByName(name)
	if err != nil {
		return nil, err
	}

	// hproc, err := windows.OpenProcess(
	// 	windows.PROCESS_VM_READ, false, pid)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (p *WinProc) Read(addr uintptr, size uint32) ([]byte, error) {
	buf := make([]byte, size)
	var read uint32 = 0

	ret, _, _ := winapiReadProcessMemory.Call(
		uintptr(p.Handle),                // hProcess
		addr,                             // lpBaseAddress
		uintptr(unsafe.Pointer(&buf[0])), // lpBuffer
		uintptr(size),                    // nSize
		uintptr(unsafe.Pointer(&read)))   // lpNumberOfBytesRead

	if ret != uintptr(windows.ERROR_SUCCESS) {
		return nil, fmt.Errorf("failed to read remote process memory: 0x%x", ret)
	}

	if read != size {
		return nil, fmt.Errorf("partial read on ReadProcessMemory at: 0x%x", addr)
	}
	return buf, nil
}
