package winutil

import (
	"encoding/binary"

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
	return ptr, err
}
