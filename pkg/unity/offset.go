package unity

import (
	"fmt"
	"reflect"
)

type Offsets struct {
	// Direct Offsets of GameObjectManager
	GameObjMgr     uintptr
	LastTaggedObj  uintptr
	FirstTaggedObj uintptr
	LastActiveObj  uintptr
	FirstActiveObj uintptr

	// Offsets from BaseObjects
	NextBaseObj    uintptr // e.g. 0x08
	GameObject     uintptr // e.g. 0x10
	GameObjectName uintptr // e.g. 0x60
	ObjectClass    uintptr // e.g. 0x30

	// Offsets from Object.ObjectClass
	Entity     uintptr // e.g. 0x18
	BaseEntity uintptr // e.g. 0x28

	// Offsets from LocalGameWorld
	PlayerListClass uintptr

	// Offsets from LocalGameWorld->PlayerList
	PlayerListObj  uintptr
	PlayerListSize uintptr

	// Offsets from PlayerList->PlayerListObj
	PlayerListData uintptr
}

func ValidateOffsetStruct(os Offsets) error {
	v := reflect.ValueOf(os)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == 0 {
			return fmt.Errorf("offset entry %d was null", i)
		}
	}
	return nil
}
