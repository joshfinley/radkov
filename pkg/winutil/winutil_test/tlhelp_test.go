package winutil_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

func TestFindProcByName(t *testing.T) {
	pid, err := winutil.FindPidByName("dlv.exe")
	if err != nil {
		t.FailNow()
	}

	t.Log(pid)
}
