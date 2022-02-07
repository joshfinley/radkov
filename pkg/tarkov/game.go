package tarkov

import (
	"encoding/binary"
	"log"
	"os"
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

var logger = log.New(os.Stdout, "TARKOV: ", 0)

func NewTarkovGame() (*unity.UnityGame, error) {
	//
	// Initialize the UnityGame, waiting as necessary
	//
	offsets := unity.Offsets{
		GameObjMgr:     0x17F8D28,
		LastTaggedObj:  0,
		FirstTaggedObj: 0x08,
		LastActiveObj:  0x20,
		FirstActiveObj: 0x28,
		NextBaseObj:    0x08,
		GameObject:     0x10,
		GameObjectName: 0x60,
		ObjectClass:    0x30,
		Entity:         0x18,
		BaseEntity:     0x28,
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
	for err == unity.ErrorGameWorldNotFound {
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

	GetAllPlayers(tg)

	return tg, nil
}

func GameMain(tg *unity.UnityGame) error {

	return nil
}

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
	for cptr := 0; cptr <= len(players); cptr = cptr + 8 {
		players[pidx] = uintptr(
			binary.LittleEndian.Uint64(
				pbuf[cptr : cptr+8]))
	}

	return nil, nil

}
