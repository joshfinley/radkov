package unity_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

func TestNewUnityGame(t *testing.T) {
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

	ug, err := unity.NewUnityGame("EscapeFromTarkov.exe", offsets)
	if err != nil {
		t.Fail()
	}

	if ug == nil {
		t.FailNow()
	}
}
