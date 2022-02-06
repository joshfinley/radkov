package unity

import "errors"

//
// GameObjMgr Types and Functions
//

type GameObjMgr uintptr

func (bg *BaseGame) FindGameObjMgr(offset uintptr) (GameObjMgr, error) {
	if bg.Proc == nil || bg.Mod == nil {
		return 0, errors.New("BaseGame info not initialized")
	}

	gom, err := bg.Proc.ReadPtr64(bg.Mod.ModuleBase + offset)
	if err != nil {
		return 0, err
	}
	if gom == 0 {
		return 0, errors.New("failed to find GameObjectManager")
	}

	return GameObjMgr(gom), nil
}

func (bg *BaseGame) GetLastTaggedObj(gom GameObjMgr) (uintptr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom))
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}

func (bg *BaseGame) GetFirstTaggedObj(gom GameObjMgr) (uintptr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom) + 0x8)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}

func (bg *BaseGame) GetLastActiveObj(gom GameObjMgr) (uintptr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom) + 0x20)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}

func (bg *BaseGame) GetFirstActiveObj(gom GameObjMgr) (uintptr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom) + 0x28)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}
