package unity

import "errors"

//
// GameObjMgr Types and Functions
//

type GameObjMgr BaseObjPtr

func (bg *BaseGame) FindGameObjMgr(offset uintptr) (GameObjMgr, error) {
	if bg.Proc == nil || bg.Mod == nil {
		return 0, errors.New("BaseGame info not initialized")
	}

	gom, err := bg.Proc.ReadPtr64(bg.Mod.ModuleBase + offset)
	if err != nil {
		return 0, err
	}

	return GameObjMgr(gom), nil
}

func (bg *BaseGame) GetLastTaggedObj(gom GameObjMgr) (BaseObjPtr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom))
	if err != nil {
		return 0, err
	}
	return BaseObjPtr(addr), nil
}

func (bg *BaseGame) GetFirstTaggedObj(gom GameObjMgr) (BaseObjPtr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom) + 0x8)
	if err != nil {
		return 0, err
	}
	return BaseObjPtr(addr), nil
}

func (bg *BaseGame) GetLastActiveObj(gom GameObjMgr) (BaseObjPtr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom) + 0x10)
	if err != nil {
		return 0, err
	}
	return BaseObjPtr(addr), nil
}

func (bg *BaseGame) GetFirstActiveObj(gom GameObjMgr) (BaseObjPtr, error) {
	addr, err := bg.Proc.ReadPtr64(uintptr(gom) + 0x18)
	if err != nil {
		return 0, err
	}
	return BaseObjPtr(addr), nil
}
