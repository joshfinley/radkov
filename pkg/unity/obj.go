package unity

import (
	"errors"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

//
// BaseGame Type and Functions
//

type BaseGame struct {
	Proc *winutil.WinProc // process associated with the game
	Mod  *winutil.WinMod  // dll associated with the game
}

func NewBaseGame(proc *winutil.WinProc) (*BaseGame, error) {
	gameMod := winutil.FindModule("UnityPlayer.dll", &proc.Modules)
	if gameMod == nil {
		return nil, errors.New("could not locate UnityPlayer.dll")
	}

	return &BaseGame{
		Proc: proc,
		Mod:  gameMod,
	}, nil
}

func (b *BaseGame) GameMain() {}

func (bg *BaseGame) GetGameObj(addr uintptr) (uintptr, error) {
	addr, err := bg.Proc.ReadPtr64(addr + 0x10)
	if err != nil {
		return 0, err
	}
	return addr, nil
}

func (bg *BaseGame) GetNextBaseObj(obj uintptr) (uintptr, error) {
	addr, err := bg.Proc.ReadPtr64(obj + 0x8)
	if err != nil {
		return 0, err
	}
	return addr, nil
}

func (bg *BaseGame) GetGameObjName(obj uintptr) (string, error) {
	nameAddr, err := bg.Proc.ReadPtr64(obj + 0x60)
	if err != nil {
		return "", err
	}
	nameBuf, err := bg.Proc.Read(nameAddr, 100)
	if err != nil {
		return "", err
	}

	return string(nameBuf), nil
}

func (bg *BaseGame) GetGameComponentAddr(obj uintptr) (uintptr, error) {
	objclass, err := bg.Proc.ReadPtr64(obj + 0x30)
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
