package tarkov

import (
	"encoding/binary"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

func GetPlayerListInstance(tg *unity.UnityGame) (uintptr, error) {
	plist, err := tg.Proc.ReadPtr64(
		tg.LocalGameWorld + tg.Offsets.PlayerListClass)
	return plist, err
}

// Get the player count from the game
func GetPlayerCount(tg *unity.UnityGame) (uint32, error) {
	instancePtr, err := GetPlayerListInstance(tg)
	if err != nil {
		return 0, err
	}
	plistSize, err := tg.Proc.ReadPtr32(
		instancePtr + tg.Offsets.PlayerListSize)
	if err != nil {
		return plistSize, err
	}
	return plistSize, err
}

// If the player list head is already found, get list size from that
// pointer
func GetPlayerCountFromListHead(tg *unity.UnityGame, plist uintptr) (uint32, error) {
	plistSize, err := tg.Proc.ReadPtr32(
		plist + tg.Offsets.PlayerListSize)
	if err != nil {
		return plistSize, err
	}
	return plistSize, err
}

func PlayerCountValid(size uint32) error {
	if size < 1 || size > 30 {
		return ErrorInvalidPlayerListSize
	}
	return nil
}

func GetPlayerListObj(tg *unity.UnityGame, plist uintptr) (uintptr, error) {
	plistObj, err := tg.Proc.ReadPtr64(
		plist + tg.Offsets.PlayerListObj)
	if err != nil {
		return 0, err
	}
	return plistObj, err
}

func GetPlayerPointers(tg *unity.UnityGame) ([]uintptr, error) {
	plist, err := GetPlayerListInstance(tg)
	if err != nil {
		return nil, err
	}

	nplayers, err := GetPlayerCountFromListHead(
		tg, plist)
	if err != nil {
		return nil, err
	}

	err = PlayerCountValid(nplayers)
	if err != nil {
		return nil, err
	}

	plistObj, err := GetPlayerListObj(tg, plist)
	if err != nil {
		return nil, err
	}

	pbuf, err := tg.Proc.Read(
		plistObj+0x20, uint32(int32(nplayers)*8))
	if err != nil {
		return nil, err
	}

	players := make([]uintptr, nplayers*8)

	pidx := 0
	for cptr := 0; cptr <= len(players) && cptr+8 <= len(players); cptr = cptr + 8 {
		players[pidx] = uintptr(
			binary.LittleEndian.Uint64(
				pbuf[cptr : cptr+8]))
		pidx++
	}

	return players, nil
}

// TODO:
// Add code to read Player Data from pointers collected by this
// Will require refactoring most of this functions code into something
// like LocalGameWorld.GetPlayerPointers() then GetTarkovPlayers([]PlayerPointers)
// func GetAllPlayers(tg *unity.UnityGame) (*[]TarkovPlayer, error) {
// 	plist, err := tg.Proc.ReadPtr64(
// 		tg.LocalGameWorld + 0x80)
// 	if err != nil {
// 		return nil, err
// 	}

// 	plistSize, err := tg.Proc.ReadPtr32(
// 		plist + 0x18)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if plistSize < 1 || plistSize > 30 {
// 		return nil, ErrorInvalidPlayerListSize
// 	}

// 	plistObj, err := tg.Proc.ReadPtr64(
// 		plist + 0x10)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pbuf, err := tg.Proc.Read(
// 		plistObj+0x20, uint32(int32(plistSize)*8))
// 	if err != nil {
// 		return nil, err
// 	}

// 	players := make([]uintptr, plistSize*8)

// 	pidx := 0
// 	for cptr := 0; cptr <= len(players) && cptr+8 <= len(players); cptr = cptr + 8 {
// 		players[pidx] = uintptr(
// 			binary.LittleEndian.Uint64(
// 				pbuf[cptr : cptr+8]))
// 	}

// 	return nil, nil
// }
