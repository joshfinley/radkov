package tarkov

import (
	"bytes"
	"encoding/binary"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

// type TarkovPlayer struct {
// 	Address uintptr    // pointer to the player object in game memory
// 	ID      string     // player's ID
// 	GroupID string     // player's group
// 	Created *time.Time // time the player object was created
// 	Scav    bool       // player is a scav
// 	Human   bool       // AI or human?
// 	//CurrentHealth int        // Unknown offsets
// 	//MaxHealth     int        // Unknown offsets
// 	LocalPosition *unity.Vec3 // player's current position
// }

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
