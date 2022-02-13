package tarkov

import (
	"log"
	"sync"
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

func MonitorGameAsync(wg sync.WaitGroup, ch chan<- unity.RawVec3) {
	return
}

func MonitorGame(offsets *unity.Offsets) error {
	tg, err := AwaitGame(offsets)
	if err != nil {
		log.Println(err)
	}

	log.Printf("EscapeFromTarkov.exe process found (pid: %d)",
		tg.Proc.Pid)
	log.Printf("UnityPlayer.dll loaded (addr: 0x%x",
		tg.Mod.ModuleBase)

	for tg.GameWorldActive() {
		//continue // placeholder until the below are completed

		// load all the players
		players, err := GetPlayerPointers(tg)
		if err != nil {
			return err
		}

		positions, err := GetPlayerPositions(tg, players)
		if err != nil {
			return err
		}
		log.Println(positions[0].Unmarshal())

		// transmit player data
		// e.g. server.publish(players)

		// check that the match is still active,
		// if not, AwaitGame()
	}

	return nil
}

// Block until the game has been launched and a match has been entered
func AwaitGame(offsets *unity.Offsets) (*unity.UnityGame, error) {
	tg, err := AwaitGameLaunch(offsets)
	if err != nil {
		if err != unity.ErrorGameWorldNotFound {
			return tg, err
		}
	}

	err = AwaitGameWorld(tg)
	if err != nil {
		return tg, err
	}

	return tg, err
}

func AwaitGameWorld(tg *unity.UnityGame) error {
	log.Println("Awaiting session start")

	var err error = unity.ErrorGameWorldNotFound
	for err == unity.ErrorGameWorldNotFound {
		err = tg.RefreshGameWorld()

		if err == nil {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	ready := CheckGameWorldReady(tg)
	for !ready {
		tg.RefreshGameWorld()
		time.Sleep(500 * time.Millisecond)
		ready = CheckGameWorldReady(tg)

	}

	log.Printf("GameWorld found (0x%x)", tg.LocalGameWorld)
	return nil
}

// Block until Game Process can be loaded
func AwaitGameLaunch(offsets *unity.Offsets) (*unity.UnityGame, error) {
	log.Println("Awaiting game startup")

	var tg *unity.UnityGame

	var err error = nil
	for tg == nil {
		tg, err = unity.NewUnityGame(
			"EscapeFromTarkov.exe",
			*offsets)
		// Dont burn up too many cycles waiting for game start
		time.Sleep(500 * time.Millisecond)
	}

	log.Printf(
		"EscapeFromTarkov.exe found (process id: %d)",
		tg.Proc.Pid)
	log.Printf(
		"Found UnityPlayer.dll (module base: 0x%x)",
		tg.Mod.ModuleBase)

	return tg, err
}

// Check if the GameWorld is ready / whether we are in a match
// Returns true if player list size is a reasonable value
// Watch out! This reads memory and doesnt really check errors
// There could be smarter ways to do this
//
// TODO: Theres a bug in here somewhere... Sometimes this reports
// the game world is ready before it truly is
func CheckGameWorldReady(tg *unity.UnityGame) bool {
	for {
		nplayers, err := GetPlayerCount(tg)
		if err != nil || nplayers < 1 || nplayers > 30 {
			return false
		} else {
			return true
		}
	}
}
