package unity

import (
	"encoding/binary"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

//
// BaseGame Type and Functions
//

type BaseGame struct {
	Proc   *winutil.WinProc // process associated with the memory location
	Addr   uintptr          // address of the memory object
	Module *winutil.WinMod  // base module of game
}

func (b *BaseGame) GameMain() {}

//
// BaseObj Type and Functions
//

// Address of an object
type BaseObjPtr uintptr

func (bg *BaseGame) GetGameObj() (GameObjPtr, error) {
	addr, err := bg.Proc.ReadPtr64(bg.Addr + 0x10)
	if err != nil {
		return 0, err
	}
	return GameObjPtr(addr), nil
}

func (bg *BaseGame) GetNextBaseObj(obj BaseObjPtr) (BaseObjPtr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(obj) + 0x8)
	if err != nil {
		return 0, err
	}
	return BaseObjPtr(addr), nil
}

//
// GameObj Type and Functions
//

type GameObjPtr uintptr

func (bg *BaseGame) GetGameObjName(obj GameObjPtr) (string, error) {
	nameAddrBuf, err := bg.Proc.Read(uintptr(obj)+0x60, 8)
	if err != nil {
		return "", err
	}
	nameAddr := binary.LittleEndian.Uint64(nameAddrBuf)
	nameBuf, err := bg.Proc.Read(uintptr(nameAddr), 100)
	if err != nil {
		return "", err
	}

	return string(nameBuf), nil
}

func (bg *BaseGame) GetGameComponentAddr(obj GameObjPtr) (uintptr, error) {
	objclass, err := bg.Proc.ReadPtr64(uintptr(obj) + 0x30)
	if err != nil {
		return 0, err
	}
	entity, err := bg.Proc.ReadPtr64(objclass + 0x18)
	if err != nil {
		return 0, err
	}
	baseEntity, err := bg.Proc.ReadPtr64(entity + 0x28)
	if err != nil {
		return 0, err
	}
	return baseEntity, nil
}
