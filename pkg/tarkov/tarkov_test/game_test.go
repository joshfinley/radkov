package tarkov_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
)

func TestGameMain(t *testing.T) {
	err := tarkov.MonitorGame(&tarkov.TarkovOffsets)
	if err != nil {
		t.FailNow()
	}

	if err != nil {
		t.FailNow()
	}
}
