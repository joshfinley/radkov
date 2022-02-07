package unity

func (ug *UnityGame) GetGameObj(addr uintptr) (uintptr, error) {
	addr, err := ug.Proc.ReadPtr64(
		addr + ug.Offsets.GameObject)
	if err != nil {
		return 0, err
	}
	return addr, nil
}

func (ug *UnityGame) GetNextBaseObj(obj uintptr) (uintptr, error) {
	addr, err := ug.Proc.ReadPtr64(
		obj + ug.Offsets.NextBaseObj)
	if err != nil {
		return 0, err
	}
	return addr, nil
}

func (ug *UnityGame) GetGameObjName(obj uintptr) (string, error) {
	nameAddr, err := ug.Proc.ReadPtr64(obj + ug.Offsets.GameObjectName)
	if err != nil {
		return "", err
	}
	nameBuf, err := ug.Proc.Read(nameAddr, 100)
	if err != nil {
		return "", err
	}

	return string(nameBuf), nil
}

func (ug *UnityGame) GetGameComponentAddr(obj uintptr) (uintptr, error) {
	objclass, err := ug.Proc.ReadPtr64(
		obj + ug.Offsets.ObjectClass)
	if err != nil {
		return 0, err
	}
	entity, err := ug.Proc.ReadPtr64(
		objclass + ug.Offsets.Entity)
	if err != nil {
		return 0, err
	}
	baseEntity, err := ug.Proc.ReadPtr64(
		entity + ug.Offsets.BaseEntity)
	if err != nil {
		return 0, err
	}
	return baseEntity, nil
}
