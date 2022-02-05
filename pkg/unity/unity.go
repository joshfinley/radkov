package unity

import (
	"errors"
	"log"
	"strings"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

//
// UnityGame Type and Functions
//

type UnityGame struct {
	BaseGame          *BaseGame  // BaseGame
	GameObjectManager GameObjMgr // GameObjectManager
	LocalGameWorld    BaseObjPtr // Local game world
}

func NewUnityGame(process string, gomOffset uintptr) (*UnityGame, error) {
	proc, err := winutil.NewWinProc(process)
	if err != nil {
		return nil, err
	}

	bg, err := NewBaseGame(proc)
	if err != nil {
		return nil, err
	}

	gom, err := bg.FindGameObjMgr(gomOffset)
	if err != nil {
		return nil, err
	}

	ug := &UnityGame{
		BaseGame:          bg,
		GameObjectManager: gom,
		LocalGameWorld:    0,
	}

	lgw, err := ug.FindLocalGameWorld()
	if err != nil {
		return nil, err
	}

	ug.LocalGameWorld = lgw
	return ug, nil
}

func (ug *UnityGame) FindLocalGameWorld() (BaseObjPtr, error) {
	activeObj, err := ug.BaseGame.GetLastActiveObj(ug.GameObjectManager)
	if err != nil {
		return 0, err
	}

	i := 0
	for uintptr(activeObj) != uintptr(ug.GameObjectManager) {
		goto loop
	inc:
		{
			//i++
			continue
		}
	loop:
		// TODO fix this hack to keep searching for GameWorld
		//time.Sleep(1 * time.Second)
		if i > 50000 {
			return 0, errors.New("GameWorld not found")
		}
		gameObj, err := ug.BaseGame.GetGameObj(uintptr(activeObj))

		if err != nil {
			goto inc
		}

		activeObjName, err := ug.BaseGame.GetGameObjName(gameObj)
		if err != nil {
			goto inc
		}

		strObjName := string(activeObjName)
		log.Println(strObjName)
		if strings.Contains(activeObjName, "GameWorld") {
			return activeObj, nil
		}

		activeObj, err = ug.BaseGame.GetNextBaseObj(activeObj)
		if err != nil {
			return 0, err
		}
		goto inc
	}

	return 0, errors.New("GameWorld not found")
}

func (ug *UnityGame) GameWorldActive() bool {
	gameWorld, err := ug.FindLocalGameWorld()
	if err == nil {
		return gameWorld != 0
	}
	return false
}
