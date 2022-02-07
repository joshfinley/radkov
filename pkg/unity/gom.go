package unity

import "errors"

//
// Functions reading offsets of GameObjectManager
//

func (ug *UnityGame) FindGameObjMgr(offset uintptr) (uintptr, error) {
	if ug.Proc == nil || ug.Mod == nil {
		return 0, errors.New("BaseGame info not initialized")
	}

	gom, err := ug.Proc.ReadPtr64(ug.Mod.ModuleBase + offset)
	if err != nil {
		return 0, err
	}
	if gom == 0 {
		return 0, errors.New("failed to find GameObjectManager")
	}

	return uintptr(gom), nil
}

func (ug *UnityGame) GetLastTaggedObj(gom uintptr) (uintptr, error) {
	addr, err := ug.Proc.ReadPtr64(gom)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}

func (ug *UnityGame) GetFirstTaggedObj(gom, offset uintptr) (uintptr, error) {
	addr, err := ug.Proc.ReadPtr64(gom + offset)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}

func (ug *UnityGame) GetLastActiveObj(gom, offset uintptr) (uintptr, error) {
	addr, err := ug.Proc.ReadPtr64(uintptr(gom) + offset)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}

func (ug *UnityGame) GetFirstActiveObj(gom, offset uintptr) (uintptr, error) {
	addr, err := ug.Proc.ReadPtr64(gom + offset)
	if err != nil {
		return 0, err
	}
	return uintptr(addr), nil
}
