package tarkov

import (
	"errors"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

type TarkovGame struct {
	UnityGame *unity.UnityGame
}

func (tg *TarkovGame) GameMain() error {
	if tg.UnityGame.Base.Proc == nil || tg.UnityGame.Base.Module == nil {
		return errors.New("tarkov process or module invalid")
	}

	gomAddr, err := tg.UnityGame.Base.Proc.ReadPtr64(
		tg.UnityGame.Base.Addr + OffsetGameObjectManager)
	if err != nil {
		return err
	}
	tg.UnityGame.GameObjectManager = unity.GameObjMgr{
		Proc: tg.UnityGame.Base.Proc,
		Addr: gomAddr,
	}
	lgw, err := tg.UnityGame.FindLocalGameWorld()
	if err != nil {
		return err
	}
	tg.UnityGame.LocalGameWorld = lgw
	return nil
}

func LoadGame() (*TarkovGame, error) {
	tkovProc, err := winutil.NewWinProc("EscapeFromTarkov.exe")
	if err != nil {
		return nil, err
	}

	gameMod := winutil.FindModule("UnityPlayer.dll", &tkovProc.Modules)

	bg := &unity.BaseGame{
		Proc:   tkovProc,
		Addr:   gameMod.ModuleBase,
		Module: gameMod,
	}

	if bg.Module == nil {
		return nil, errors.New("failed to locate UnityPlayer.dll in Tarkov process")
	}

	ug := &unity.UnityGame{
		Base: bg,
		GameObjectManager: unity.GameObjMgr{
			Proc: tkovProc,
			Addr: 0,
		},
		LocalGameWorld: &unity.BaseObj{
			Proc: tkovProc,
			Addr: 0,
		},
	}

	tg := TarkovGame{
		UnityGame: ug,
	}

	return &tg, nil
}
