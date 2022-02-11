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

func TestGetTarkovPlayer(t *testing.T) {
	tg, err := tarkov.AwaitGame(&tarkov.TarkovOffsets)
	if err != nil {
		t.FailNow()
	}

	plist, err := tarkov.GetPlayerListBuffer(tg)
	if err != nil {
		t.FailNow()
	} else if len(plist) < 1 {
		t.FailNow()
	}

	for _, player := range plist {
		tarkov.GetTarkovPlayer(tg, player)
		t.Log(player)
	}

}
