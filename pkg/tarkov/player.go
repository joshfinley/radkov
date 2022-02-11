package tarkov

import (
	"bytes"
	"encoding/binary"
	"log"
	"strings"
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

type TarkovPlayer struct {
	Address uintptr    // pointer to the player object in game memory
	ID      string     // player's ID
	GroupID string     // player's group
	Created *time.Time // time the player object was created
	Scav    bool       // player is a scav
	Human   bool       // AI or human?
	//CurrentHealth int        // Unknown offsets
	//MaxHealth     int        // Unknown offsets
	LocalPosition *unity.Vec3 // player's current position
}

func GetPlayerIdString(tg *unity.UnityGame, playerPtr uintptr) (string, error) {
	return tg.ReadUnityEngineString(
		playerPtr + tg.Offsets.PlayerProfilePlayerID)
}

func GetPlayerInfoAddr(tg *unity.UnityGame, playerPtr uintptr) (uintptr, error) {
	profilePtr, err := tg.Proc.ReadPtr64(
		playerPtr + tg.Offsets.PlayerProfile)
	if err != nil {
		return profilePtr, err
	}

	return profilePtr, nil
}

func GetPlayerInfoGroupId(tg *unity.UnityGame, playerInfoAddr uintptr) (string, error) {
	return tg.ReadUnityEngineString(
		playerInfoAddr + tg.Offsets.PlayerProfilePlayerID)
}

func GetPlayerInfoName(tg *unity.UnityGame, playerInfoAddr uintptr) (string, error) {
	return tg.ReadUnityEngineString(
		playerInfoAddr + tg.Offsets.PlayerInfoName)
}

func GetPlayerInfoCreationTime(tg *unity.UnityGame, playerInfoAddr uintptr) (*time.Time, error) {
	t, err := tg.Proc.ReadPtr32(
		playerInfoAddr + tg.Offsets.PlayerInfoCreationTime)
	if err != nil {
		return nil, err
	}
	tint := int64(t)
	tres := time.Unix(tint, 0)

	return &tres, nil
}

func PlayerInfoIsLocal(tg *unity.UnityGame, playerInfoPtr uintptr) bool {
	// were doing enough error handling - the result of this function
	// shouldnt be consequential enough to warrant it here too
	ret, _ := tg.Proc.ReadPtr64(playerInfoPtr + tg.Offsets.PlayerIsLocal)
	return ret == 0
}

func PlayerInfoIsScav(tg *unity.UnityGame, playerInfoAddr uintptr) bool {
	// Not sure how well this will work
	t, _ := tg.Proc.ReadPtr32(
		playerInfoAddr + tg.Offsets.PlayerInfoCreationTime)

	return t == 0
}

func PlayerInfoHumanScav(tg *unity.UnityGame, playerInfo uintptr) bool {
	name, _ := GetPlayerInfoName(tg, playerInfo)
	return !PlayerInfoIsScav(tg, playerInfo) && strings.Contains(name, " ")
}

func GetPlayerLocalPosition(tg *unity.UnityGame, player uintptr) (*unity.Vec3, error) {
	ctx, err := tg.Proc.ReadPtr64(player + tg.Offsets.PlayerMovementCtx)
	if err != nil {
		return nil, err
	}
	posData, err := tg.Proc.Read(ctx+tg.Offsets.MvmtCtxLocalPos, 8*3)
	var vec unity.Vec3
	posBytes := bytes.NewBuffer(posData)
	binary.Read(posBytes, binary.LittleEndian, vec)
	return &vec, nil
}

func GetTarkovPlayer(tg *unity.UnityGame, playerPtr uintptr) (*TarkovPlayer, error) {
	playerInfoAddr, err := GetPlayerInfoAddr(tg, playerPtr)
	if err != nil {
		log.Println(err)
	}
	if err == nil {
		log.Println("nil error!")
	}

	pidString, err := GetPlayerIdString(tg, playerInfoAddr)
	if err != nil {
		log.Println(err)
	}

	gidString, err := GetPlayerInfoGroupId(tg, playerInfoAddr)
	if err != nil {
		log.Println(err)
	}

	pcreated, err := GetPlayerInfoCreationTime(tg, playerInfoAddr)
	if err != nil {
		log.Println(err)
	}

	isScav := PlayerInfoIsScav(tg, playerInfoAddr)
	pos, err := GetPlayerLocalPosition(tg, playerPtr)
	if err != nil {
		log.Println(err)
	}

	player := &TarkovPlayer{
		Address:       playerPtr,
		ID:            pidString,
		GroupID:       gidString,
		Created:       pcreated,
		Scav:          isScav,
		Human:         false,
		LocalPosition: pos,
	}

	return player, nil
}
