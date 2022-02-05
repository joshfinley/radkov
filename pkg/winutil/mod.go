package winutil

import (
	"golang.org/x/sys/windows"
)

type WinMod struct {
	ModuleBase  uintptr // module base address
	SizeOfImage uint32  // module size
	FullDllName string  // module name
}

func GetProcModules(hproc windows.Handle) (*[]WinMod, error) {
	hmods := make([]windows.Handle, 1024)
	winmods := make([]WinMod, 0)
	var cbn uint32 = 0

	windows.EnumProcessModules(hproc, &hmods[0], 1024, &cbn)
	cbn = cbn / 8
	hmods = append([]windows.Handle(nil), hmods[:cbn]...)
	namebuf := make([]uint16, 256*2)
	for i := 0; i < int(cbn); i++ {
		// _, _, err := winapiGetModuleFileNameEx.Call(
		// 	uintptr(hproc),
		// 	uintptr(hmods[i]),
		// 	uintptr(unsafe.Pointer(&filename[0])),
		// 	256)

		err := windows.GetModuleFileNameEx(
			hproc,
			hmods[i],
			&namebuf[0],
			256*2)

		if err != nil {
			return nil, err
		}

		modinfo := windows.ModuleInfo{}
		err = windows.GetModuleInformation(
			hproc,
			hmods[i],
			&modinfo,
			24) // if this doesnt work try 24
		if err != nil {
			return nil, err
		}

		// convert stupid wide char buffer to go string
		filename := windows.UTF16ToString(namebuf)

		winmods = append(winmods, WinMod{
			ModuleBase:  modinfo.BaseOfDll,
			SizeOfImage: modinfo.SizeOfImage,
			FullDllName: string([]byte(filename)),
		})
	}

	return &winmods, nil
}
