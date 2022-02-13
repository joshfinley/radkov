package tarkov

import "gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"

//
// GameOffsets, subject to change.
//
var TarkovOffsets = unity.Offsets{
	GameObjMgr:              0x17F8D28,
	LastTaggedObj:           0,
	FirstTaggedObj:          0x08,
	LastActiveObj:           0x20,
	FirstActiveObj:          0x28,
	NextBaseObj:             0x08,
	GameObject:              0x10,
	GameObjectName:          0x60,
	ObjectClass:             0x30,
	Entity:                  0x18,
	BaseEntity:              0x28,
	PlayerListClass:         0x80,
	PlayerListObj:           0x10,
	PlayerListSize:          0x18,
	PlayerListData:          0x20,
	PlayerIsLocal:           0x18,
	PlayerProfile:           0x4B8,
	PlayerBody:              0xa8,
	PlayerMovementCtx:       0x40,
	PlayerHandsController:   0x488,
	PlayerHealth:            0x470,
	MvmtCtxLocalPos:         0x23c,
	PlayerProfilePlayerID:   0x10,
	PlayerProfilePlayerInfo: 0x28,
	PlayerInfoName:          0x10,
	PlayerInfoGroupID:       0x18,
	PlayerInfoCreationTime:  0x54,
	PlayerInfoAcctType:      0x60,
	EngineStringSize:        0x10,
	EngineStringData:        0x14,
}
