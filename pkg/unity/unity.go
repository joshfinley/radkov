package unity

import (
	"errors"
	"strings"
)

//
// UnityGame Type and Functions
//

type UnityGame struct {
	BaseGame          *BaseGame   // BaseGame
	GameObjectManager GameObjMgr  // GameObjectManager
	LocalGameWorld    *BaseObjPtr // Local game world
}

func (ug *UnityGame) FindLocalGameWorld() (BaseObjPtr, error) {
	activeObj, err := ug.BaseGame.GetFirstActiveObj(ug.GameObjectManager)
	if err != nil {
		return 0, err
	}

	i := 0
	for uintptr(activeObj) != uintptr(ug.GameObjectManager) {
		if i > 50000 {
			return 0, errors.New("GameWorld not found")
		}
		gameObj, err := ug.BaseGame.GetGameObj()

		if err != nil {
			return 0, err
		}

		activeObjName, err := ug.BaseGame.GetGameObjName(gameObj)
		if err != nil {
			return 0, err
		}

		if strings.Contains(activeObjName, "GameWorld") {
			return activeObj, nil
		}

		activeObj, err = ug.BaseGame.GetNextBaseObj(activeObj)
		if err != nil {
			return 0, err
		}
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
