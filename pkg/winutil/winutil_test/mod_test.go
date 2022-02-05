package winutil_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
	"golang.org/x/sys/windows"
)

func TestGetProcModules(t *testing.T) {
	pid, err := winutil.FindPidByName("dlv.exe")
	if err != nil {
		t.FailNow()
	}

	t.Log(pid)
	hproc, err := windows.OpenProcess(
		winutil.PROCESS_ALL_ACCESS,
		false,
		pid)

	if err != nil {
		t.FailNow()
	}

	mods, err := winutil.GetProcModules(hproc)
	if err != nil {
		t.FailNow()
	}

	t.Log(mods)
}
