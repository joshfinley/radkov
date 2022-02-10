package tarkov

import (
	"encoding/binary"
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

type TarkovPlayer struct {
	Address       uintptr    // pointer to the player object in game memory
	ID            string     // player's ID
	GroupID       string     // player's group
	Created       time.Time  // time the player object was created
	Scav          bool       // player is a scav
	Human         bool       // AI or human?
	CurrentHealth int        //
	MaxHealth     int        //
	LocalPosition unity.Vec3 // player's current position
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
