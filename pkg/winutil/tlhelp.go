package winutil

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type tlproc struct {
	pid  int
	ppid int
	exe  string
}

func getProcSnapshot() ([]tlproc, error) {
	hsnap, err := windows.CreateToolhelp32Snapshot(
		TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(hsnap)
	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	err = windows.Process32First(hsnap, &entry)
	if err != nil {
		return nil, err
	}

	res := make([]tlproc, 0)
	for {
		res = append(res, newTlproc(&entry))
		err = windows.Process32Next(hsnap, &entry)
		if err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return res, nil
			}
			return nil, err
		}
	}
}

func newTlproc(e *windows.ProcessEntry32) tlproc {
	// Find when the string ends for decoding
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return tlproc{
		pid:  int(e.ProcessID),
		ppid: int(e.ParentProcessID),
		exe:  syscall.UTF16ToString(e.ExeFile[:end]),
	}
}

func FindPidByName(name string) (uint32, error) {
	procs, err := getProcSnapshot()
	if err != nil {
		return 0, err
	}
	for _, p := range procs {
		if strings.EqualFold(p.exe, name) {
			return uint32(p.pid), nil
		}
	}
	return 0, fmt.Errorf("could not find process '%s'", name)
}
