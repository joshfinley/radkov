package winutil

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/sys/windows"
)

type WinProc struct {
	Pid     uint32         // process id
	Modules []WinMod       // loaded dlls
	Handle  windows.Handle // handle to process
}

func NewWinProc(name string) (*WinProc, error) {
	pid, err := FindPidByName(name)
	if err != nil {
		return nil, err
	}

	hproc, err := windows.OpenProcess(
		PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return nil, err
	}

	mods, err := GetProcModules(hproc)
	if err != nil {
		return nil, err
	}

	return &WinProc{
		Pid:     pid,
		Modules: *mods,
		Handle:  hproc,
	}, nil
}

func (p *WinProc) Read(addr uintptr, size uint32) ([]byte, error) {
	buf := make([]byte, size)
	var read uintptr = 0
	err := windows.ReadProcessMemory(
		p.Handle,
		addr,
		&buf[0],
		uintptr(size),
		&read)
	if err != nil {
		return nil, err
	}

	return buf, err
}

// Read 64 bits (8 byte) at addr
func (p *WinProc) ReadPtr64(addr uintptr) (uintptr, error) {
	buf := make([]byte, 8)
	var read uintptr = 0
	err := windows.ReadProcessMemory(
		p.Handle,
		addr,
		&buf[0],
		uintptr(8),
		&read)
	if err != nil {
		return 0, err
	}
	ptr := uintptr(binary.LittleEndian.Uint64(buf))
	if ptr == 0 {
		return 0, fmt.Errorf("failed to read memory at 0x%x", addr)
	}
	return ptr, err
}

// Read 32 bits (4 bytes) at addr
func (p *WinProc) ReadPtr32(addr uintptr) (uint32, error) {
	buf := make([]byte, 4)
	var read uintptr = 0
	err := windows.ReadProcessMemory(
		p.Handle,
		addr,
		&buf[0],
		uintptr(4),
		&read)
	if err != nil {
		return 0, err
	}
	val := uint32(binary.LittleEndian.Uint32(buf))
	if val == 0 {
		return 0, fmt.Errorf("failed to read memory at 0x%x", addr)
	}

	return val, err
}
