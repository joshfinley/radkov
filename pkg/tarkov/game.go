package tarkov

import (
	"log"
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

//
// GameOffsets, subject to change.
//
var TarkovOffsets = unity.Offsets{
	GameObjMgr:      0x17F8D28,
	LastTaggedObj:   0,
	FirstTaggedObj:  0x08,
	LastActiveObj:   0x20,
	FirstActiveObj:  0x28,
	NextBaseObj:     0x08,
	GameObject:      0x10,
	GameObjectName:  0x60,
	ObjectClass:     0x30,
	Entity:          0x18,
	BaseEntity:      0x28,
	PlayerListClass: 0x80,
	PlayerListObj:   0x10,
	PlayerListData:  0x20,
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

	for {
		continue // placeholder until the below are completed

		// load all the players
		// e.g. players == GetAllPlayers()

		// transmit player data
		// e.g. server.publish(players)

		// check that the match is still active,
		// if not, AwaitGame()
	}
}

// Block until the game has been launched and a match has been entered
func AwaitGame(offsets *unity.Offsets) (*unity.UnityGame, error) {
	tg, err := AwaitGameLaunch(offsets)
	if err != nil {
		return tg, err
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
		ready := CheckGameWorldReady(tg)
		if err == nil && ready == true {
			return nil
		}
	}

	return err
}

// Block until Game Process can be loaded
func AwaitGameLaunch(offsets *unity.Offsets) (*unity.UnityGame, error) {
	log.Println("Awaiting game startup")

	tg := &unity.UnityGame{
		Proc:              nil,
		Mod:               nil,
		GameObjectManager: 0,
		LocalGameWorld:    0,
		Offsets:           *offsets,
	}

	var err error = nil
	for tg.Proc == nil {
		tg, err = unity.NewUnityGame(
			"EscapeFromTarkov.exe",
			*offsets)
		// Dont burn up too many cycles waiting for game start
		time.Sleep(500 * time.Millisecond)
	}

	return tg, err
}

// Check if the GameWorld is ready / whether we are in a match
// Returns true if player list size is a reasonable value
// Watch out! This reads memory and doesnt really check errors
// There could be smarter ways to do this
func CheckGameWorldReady(tg *unity.UnityGame) bool {
	for {
		nplayers, err := GetPlayerListSize(tg)
		if err != nil || nplayers < 1 || nplayers > 30 {
			// Dont burn too many cycles waiting for match load
			time.Sleep(500 * time.Millisecond)
		} else {
			return true
		}
	}
}
