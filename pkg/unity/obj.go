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

type BaseObj struct {
	Proc *winutil.WinProc // process associated with the memory location
	Addr uintptr          // address of the memory object
}

func (obj *BaseObj) GetGameObj() (*GameObj, error) {
	addr, err := obj.Proc.ReadPtr64(obj.Addr + 0x10)
	if err != nil {
		return nil, err
	}
	return &GameObj{
		Proc: obj.Proc,
		Addr: addr,
	}, nil
}

func (obj *BaseObj) GetNextBaseObj() (*BaseObj, error) {
	addr, err := obj.Proc.ReadPtr64(obj.Addr + 0x8)
	if err != nil {
		return nil, err
	}
	return &BaseObj{
		Proc: obj.Proc,
		Addr: addr,
	}, nil
}

//
// GameObj Type and Functions
//

type GameObj BaseObj

func (obj *GameObj) GetGameObjName() (string, error) {
	nameAddrBuf, err := obj.Proc.Read(obj.Addr+0x60, 8)
	if err != nil {
		return "", err
	}
	nameAddr := binary.LittleEndian.Uint64(nameAddrBuf)
	nameBuf, err := obj.Proc.Read(uintptr(nameAddr), 100)
	if err != nil {
		return "", err
	}

	return string(nameBuf), nil
}

func (obj *GameObj) GetGameComponentAddr() (uintptr, error) {
	objclass, err := obj.Proc.ReadPtr64(obj.Addr + 0x30)
	if err != nil {
		return 0, err
	}
	entity, err := obj.Proc.ReadPtr64(objclass + 0x18)
	if err != nil {
		return 0, err
	}
	baseEntity, err := obj.Proc.ReadPtr64(entity + 0x28)
	if err != nil {
		return 0, err
	}
	return baseEntity, nil
}
