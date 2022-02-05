package unity

//
// GameObjMgr Types and Functions
//

type GameObjMgr BaseObjPtr

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
