package unity

import (
	"errors"
	"strings"
)

//
// UnityGame Type and Functions
//

type UnityGame struct {
	Base              *BaseGame  // BaseGame
	GameObjectManager GameObjMgr // GameObjectManager
	LocalGameWorld    *BaseObj   // Local game world
}

func (ug *UnityGame) FindLocalGameWorld() (*BaseObj, error) {
	activeObj, err := ug.GameObjectManager.GetFirstActiveObj()
	if err != nil {
		return nil, err
	}

	i := 0
	for activeObj.Addr != ug.GameObjectManager.Addr {
		if i > 50000 {
			return nil, errors.New("GameWorld not found")
		}

		gameObj, err := activeObj.GetGameObj()
		if err != nil {
			return nil, err
		}

		activeObjName, err := gameObj.GetGameObjName()
		if err != nil {
			return nil, err
		}

		if strings.Contains(activeObjName, "GameWorld") {
			return activeObj, nil
		}

		activeObj, err = activeObj.GetNextBaseObj()
		if err != nil {
			return nil, err
		}
	}

	return nil, errors.New("GameWorld not found")
}

func (ug *UnityGame) GameWorldActive() bool {
	gameWorld, err := ug.FindLocalGameWorld()
	if err == nil {
		return gameWorld != nil
	}
	return false
}
