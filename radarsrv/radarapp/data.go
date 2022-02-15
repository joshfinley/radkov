package radarapp

import (
	"errors"
	"sync"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
	"golang.org/x/sync/errgroup"
)

// Global game state data (simple route for serving data in HTTP application)
// Must be initialized by calling application (package main)
// TODO: make application init this on startup
var GlobalGameState GameState

// Thread safe game state type
type GameState struct {
	Mutex           *sync.RWMutex
	PlayerPositions []unity.Vec2
}

func (gs *GameState) Init() {
	gs.Mutex = &sync.RWMutex{}
	gs.PlayerPositions = []unity.Vec2{}
}

func (gs *GameState) IsInitialized() bool {
	if gs.PlayerPositions == nil || gs.Mutex == nil {
		return false
	}

	return true
}

func (gs *GameState) GetPlayerPositions() (*[]unity.Vec2, error) {
	if !gs.IsInitialized() {
		return nil, errors.New("game state not initialized")
	}
	return &gs.PlayerPositions, nil
}

func (gs *GameState) SetPlayerPositions(new [][]byte) error {
	// lock the mutex
	gs.Mutex.Lock()
	defer gs.Mutex.Unlock()

	// clear out old data
	gs.PlayerPositions = make([]unity.Vec2, len(new))

	// set the new data
	var r errgroup.Group
	for i, pos := range new {
		i, pos := i, pos
		r.Go(func() error {
			res := unity.UnmarshalVec2(pos)
			gs.PlayerPositions[i] = res
			return nil
		})
	}

	if err := r.Wait(); err != nil {
		return err
	}
	return nil
}
