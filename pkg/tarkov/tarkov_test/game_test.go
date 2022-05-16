package tarkov_test

import (
	"fmt"
	"testing"

	"github.com/joshfinley/radkov/pkg/tarkov"
	"github.com/joshfinley/radkov/pkg/unity"
)

func TestMonitorGame(t *testing.T) {
	pch := make(chan [][]byte)
	go tarkov.MonitorGame(pch, &tarkov.TarkovOffsets)

	for d := range pch {
		fmt.Println(unity.UnmarshalVec2(d[0]))
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
		t.Log(unity.UnmarshalVec2(positions[0]))
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
