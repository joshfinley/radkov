package tarkov_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
)

func TestMonitorGame(t *testing.T) {
	tg, err := tarkov.AwaitGame(&tarkov.TarkovOffsets)
	if err != nil {
		t.FailNow()
	}

	t.Log("EscapeFromTarkov.exe process found (pid: %d)",
		tg.Proc.Pid)
	t.Log("UnityPlayer.dll loaded (addr: 0x%x",
		tg.Mod.ModuleBase)

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

		// transmit player data
		// e.g. server.publish(players)

		// check that the match is still active,
		// if not, AwaitGame()
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
