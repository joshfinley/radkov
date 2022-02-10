package unity

import (
	"errors"
	"strings"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/winutil"
)

//
// UnityGame Type and Functions
//

type UnityGame struct {
	Proc              *winutil.WinProc // process associated with the game
	Mod               *winutil.WinMod  // dll associated with the game
	GameObjectManager uintptr          // GameObjectManager address
	LocalGameWorld    uintptr          // Local game world
	Offsets           Offsets
}

func NewUnityGame(process string, offsets Offsets) (*UnityGame, error) {
	err := ValidateOffsetStruct(offsets)
	if err != nil {
		return nil, err
	}

	proc, err := winutil.NewWinProc(process)
	if err != nil {
		return nil, err
	}

	gameMod := winutil.FindModule("UnityPlayer.dll", &proc.Modules)
	if gameMod == nil {
		return nil, errors.New("could not locate UnityPlayer.dll")
	}

	ug := &UnityGame{
		Proc:              proc,
		Mod:               gameMod,
		GameObjectManager: 0,
		LocalGameWorld:    0,
		Offsets:           offsets,
	}

	gom, err := ug.FindGameObjMgr(offsets.GameObjMgr)
	if err != nil {
		return nil, err
	}
	ug.GameObjectManager = gom

	lgw, err := ug.FindLocalGameWorld()
	if err != nil {
		ug.LocalGameWorld = 0
		return ug, err
	}
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

// Advance the ug.LocalGameWorld to the next found GameWorld
func (ug *UnityGame) NextGameWorld() error {
	activeObj, err := ug.GetFirstActiveObj(
		ug.GameObjectManager, ug.Offsets.FirstActiveObj)
	if err != nil {
		return err
	}

	i := 0
	for uintptr(activeObj) != uintptr(ug.GameObjectManager) {
		goto loop
	inc:
		{
			i++
			activeObj, err = ug.GetNextBaseObj(activeObj)
			if err != nil {
				return err
			}
			continue
		}
	loop:
		// Prevent this from blocking if GameWorld not found
		if i > 50000 {
			return ErrorGameWorldNotFound
		}
		gameObj, err := ug.GetGameObj(uintptr(activeObj))
		if err != nil {
			goto inc
		}

		activeObjName, err := ug.GetGameObjName(gameObj)
		if err != nil {
			goto inc
		}

		//strObjName := string(activeObjName)
		if strings.Contains(activeObjName, "GameWorld") {
			nextObj, err := ug.GetGameComponentAddr(gameObj)
			if err == nil {
				return err
			}

			if nextObj != ug.LocalGameWorld {
				ug.LocalGameWorld = nextObj
				return nil
			}
		}
		goto inc
	}

	return errors.New("GameWorld not found")
}

// Get the GameWorld object from the GameObjectManager
// Thanks to:
// https://www.unknowncheats.me/forum/escape-from-tarkov/226519-escape-tarkov-reversal-structs-offsets-310.html#post3353153
// See #6195
func (ug *UnityGame) FindLocalGameWorld() (uintptr, error) {
	activeObj, err := ug.GetFirstActiveObj(
		ug.GameObjectManager, ug.Offsets.FirstActiveObj)
	if err != nil {
		return 0, err
	}

	i := 0
	for uintptr(activeObj) != uintptr(ug.GameObjectManager) {
		goto loop
	inc:
		{
			i++
			activeObj, err = ug.GetNextBaseObj(activeObj)
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
			return 0, ErrorGameWorldNotFound
		}
		gameObj, err := ug.GetGameObj(uintptr(activeObj))
		if err != nil {
			goto inc
		}

		activeObjName, err := ug.GetGameObjName(gameObj)
		if err != nil {
			goto inc
		}

		//strObjName := string(activeObjName)
		if strings.Contains(activeObjName, "GameWorld") {
			return ug.GetGameComponentAddr(gameObj)
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
