package tarkov_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
)

func TestMonitorGame(t *testing.T) {
	err := tarkov.MonitorGame(&tarkov.TarkovOffsets)
	if err != nil {
		t.FailNow()
	}

	if err != nil {
		t.FailNow()
	}
}

func TestGetPlayerPositions(t *testing.T) {
	tg, err := tarkov.AwaitGame(&tarkov.TarkovOffsets)
	if err != nil {
		t.FailNow()
	}

	players, err := tarkov.GetPlayerPointers(tg)
	if err != nil {
		t.FailNow()
	}

	positions, err := tarkov.GetPlayerPositions(tg, players)
	if err != nil {
		t.FailNow()
	}
	t.Log(positions)
}
