package tarkov

import (
	"encoding/binary"
	"log"
	"os"
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
	"golang.org/x/sys/windows"
)

var logger = log.New(os.Stdout, "TARKOV: ", 0)

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

// TODO:
// Refactor this function. Need something that can just run
// all the time and identify when an actual GameWorld has been created
func NewTarkovGame2() (*unity.UnityGame, error) {
	//
	// Initialize the UnityGame, waiting as necessary
	//
	offsets := unity.Offsets{
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
	tg, _ := unity.NewUnityGame(
		"EscapeFromTarkov.exe", offsets)

	for tg.Proc == nil {
		tg, _ = unity.NewUnityGame(
			"EscapeFromTarkov.exe",
			offsets)
		// Dont burn up too many cycles waiting for game start
		time.Sleep(500 * time.Millisecond)
	}
	logger.Println(
		"Found 'EscapeFromTarkov.exe'. Process ID: ",
		tg.Proc.Pid)

	//
	// Wait for the GameWorld
	//
	err := unity.ErrorGameWorldNotFound
	start := time.Now()
	var elapsed time.Duration
	for err == unity.ErrorGameWorldNotFound || tg.LocalGameWorld == 0 {
		err = tg.RefreshGameWorld()
		// Dont burn up too many cycles waiting for game world
		time.Sleep(500 * time.Millisecond)
		elapsed = time.Since(start)
		logger.Printf(
			"Waiting on GameWorld (%d seconds elapsed)...",
			int(elapsed.Seconds()))
	}
	logger.Printf(
		"Found GameWorld: 0x%x", tg.LocalGameWorld)

	//
	// Find active players
	//
	players, err := GetAllPlayers(tg)
	for err != nil {
		if err == ErrorInvalidPlayerListSize {
			goto next_gw
		}
		if err == windows.ERROR_PARTIAL_COPY {
			goto next_gw
		}
		if err == windows.ERROR_NOACCESS {
			goto next_gw
		}
	next_gw:
		tg.NextGameWorld()
		logger.Printf("Next game world: 0x%x", tg.LocalGameWorld)
		players, err = GetAllPlayers(tg)
	}

	log.Println(players)

	return tg, nil
}

func GameMain(tg *unity.UnityGame) error {

	return nil
}

func GetPlayerListClassPtr(tg *unity.UnityGame) (uintptr, error) {
	plist, err := tg.Proc.ReadPtr64(
		tg.LocalGameWorld + tg.Offsets.PlayerListClass)
	return plist, err
}

func GetPlayerListSize(tg *unity.UnityGame) (uint32, error) {
	instancePtr, err := GetPlayerListClassPtr(tg)
	if err != nil {
		return 0, err
	}
	plistSize, err := tg.Proc.ReadPtr32(
		instancePtr + 0x18)
	if err != nil {
		return plistSize, err
	}
	return plistSize, err
}

// TODO:
// Add code to read Player Data from pointers collected by this
// Will require refactoring most of this functions code into something
// like LocalGameWorld.GetPlayerPointers() then GetTarkovPlayers([]PlayerPointers)
func GetAllPlayers(tg *unity.UnityGame) (*[]TarkovPlayer, error) {
	plist, err := tg.Proc.ReadPtr64(
		tg.LocalGameWorld + 0x80)
	if err != nil {
		return nil, err
	}

	plistSize, err := tg.Proc.ReadPtr32(
		plist + 0x18)
	if err != nil {
		return nil, err
	}

	if plistSize < 1 || plistSize > 30 {
		return nil, ErrorInvalidPlayerListSize
	}

	plistObj, err := tg.Proc.ReadPtr64(
		plist + 0x10)
	if err != nil {
		return nil, err
	}

	pbuf, err := tg.Proc.Read(
		plistObj+0x20, uint32(int32(plistSize)*8))
	if err != nil {
		return nil, err
	}

	players := make([]uintptr, plistSize*8)

	pidx := 0
	for cptr := 0; cptr <= len(players) && cptr+8 <= len(players); cptr = cptr + 8 {
		players[pidx] = uintptr(
			binary.LittleEndian.Uint64(
				pbuf[cptr : cptr+8]))
	}

	return nil, nil
}
