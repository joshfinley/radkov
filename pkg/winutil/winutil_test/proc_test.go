package winutil_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

func TestNewWinProc(t *testing.T) {
	winp, err := winutil.NewWinProc("dlv.exe")
	if err != nil {
		t.FailNow()
	}

	t.Log(winp)
}

func TestWinProcRead(t *testing.T) {
	winp, err := winutil.NewWinProc("dlv.exe")
	if err != nil {
		t.FailNow()
	}

	buf, err := winp.Read(winp.Modules[0].ModuleBase, 2)
	if err != nil {
		t.FailNow()
	}

	if string(buf) != "MZ" {
		t.FailNow()
	}
}
