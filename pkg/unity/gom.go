package unity

//
// GameObjMgr Types and Functions
//

type GameObjMgr BaseObj

func (gom *GameObjMgr) GetLastTaggedObj() (*BaseObj, error) {
	addr, err := gom.Proc.ReadPtr64(gom.Addr)
	if err != nil {
		return nil, err
	}
	return &BaseObj{
		Proc: gom.Proc,
		Addr: addr,
	}, nil
}

func (gom *GameObjMgr) GetFirstTaggedObj() (*BaseObj, error) {
	addr, err := gom.Proc.ReadPtr64(gom.Addr + 0x8)
	if err != nil {
		return nil, err
	}
	return &BaseObj{
		Proc: gom.Proc,
		Addr: addr,
	}, nil
}

func (gom *GameObjMgr) GetLastActiveObj() (*BaseObj, error) {
	addr, err := gom.Proc.ReadPtr64(gom.Addr + 0x10)
	if err != nil {
		return nil, err
	}
	return &BaseObj{
		Proc: gom.Proc,
		Addr: addr,
	}, nil
}

func (gom *GameObjMgr) GetFirstActiveObj() (*BaseObj, error) {
	addr, err := gom.Proc.ReadPtr64(gom.Addr + 0x18)
	if err != nil {
		return nil, err
	}
	return &BaseObj{
		Proc: gom.Proc,
		Addr: addr,
	}, nil
}
