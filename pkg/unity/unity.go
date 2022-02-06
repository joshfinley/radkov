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
	LocalGameWorld    uintptr    // Local game world
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
	log.Printf(
		"Game process found. UnityPlayer.dll base:0x%x",
		bg.Mod.ModuleBase)

	gom, err := bg.FindGameObjMgr(gomOffset)
	if err != nil {
		return nil, err
	}
	log.Printf(
		"GameObjectManager found:UnityPlayer.dll+0x%x",
		gomOffset)

	ug := &UnityGame{
		BaseGame:          bg,
		GameObjectManager: gom,
		LocalGameWorld:    0,
	}

	lgw, err := ug.FindLocalGameWorld()
	if err != nil {
		return nil, err
	}
	log.Printf("Local Game World found:0x%x", lgw)
	ug.LocalGameWorld = lgw
	return ug, nil
}

// Update the pointer to the GameWorld object
func (ug *UnityGame) RefreshGameWorld() error {
	gw, err := ug.FindLocalGameWorld()
	if err != nil {
		return err
	}
	ug.LocalGameWorld = gw
	return nil
}

// Get the GameWorld object from the GameObjectManager
// Thanks to:
// https://www.unknowncheats.me/forum/escape-from-tarkov/226519-escape-tarkov-reversal-structs-offsets-310.html#post3353153
// See #6195
func (ug *UnityGame) FindLocalGameWorld() (uintptr, error) {
	activeObj, err := ug.BaseGame.GetFirstActiveObj(
		ug.GameObjectManager)
	if err != nil {
		return 0, err
	}

	i := 0
	for uintptr(activeObj) != uintptr(ug.GameObjectManager) {
		goto loop
	inc:
		{
			//i++
			activeObj, err = ug.BaseGame.GetNextBaseObj(activeObj)
			if err != nil {
				return 0, err
			}
			continue
		}
	loop:
		// TODO
		// Ensure that the local game world is the correct one...
		// Maybe not the responsibility of this function?
		// Below is a hack to make sure we dont look terribly too
		// far...
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

		//strObjName := string(activeObjName)
		if strings.Contains(activeObjName, "GameWorld") {
			return activeObj, nil
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
