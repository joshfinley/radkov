package tarkov_test

import (
	"fmt"
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

func TestMonitorGame(t *testing.T) {
	pch := make(chan []unity.RawVec2)
	go tarkov.MonitorGame(pch, &tarkov.TarkovOffsets)

	for d := range pch {
		fmt.Println(d[0].Unmarshal())
	}
}

func TestMonitorGame2(t *testing.T) {
	tg, err := tarkov.AwaitGame(&tarkov.TarkovOffsets)
	if err != nil {
		t.FailNow()
	}

	// get the initial list of players
	players, err := tarkov.GetPlayerPointers(tg)
	if err != nil {
		t.FailNow()
	}

	if players == nil {
		tg, err = tarkov.AwaitGame(
			&tarkov.TarkovOffsets)
		if err != nil {
			t.FailNow()
		}
	}

	for {
		if !tg.GameWorldActive() {
			tg, err = tarkov.AwaitGame(
				&tarkov.TarkovOffsets) // if the game world goes inactive, restart the wait
			if err != nil {
				t.Log(err)
			}
		}

		// load all the players
		players, err = tarkov.GetPlayerPointers(tg)
		if err != nil {
			t.FailNow()
		}

		positions, err := tarkov.GetPlayerPositions(
			tg, players)
		if err != nil {
			t.FailNow()
		}
		t.Log(positions[0].Unmarshal())

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
