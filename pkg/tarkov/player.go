package tarkov

import (
	"time"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

type TarkovPlayer struct {
	Address       uintptr    // pointer to the player object in game memory
	ID            string     // player's ID
	GroupID       string     // player's group
	Created       time.Time  // time the player object was created
	Scav          bool       // player is a scav
	Human         bool       // AI or human?
	CurrentHealth int        //
	MaxHealth     int        //
	LocalPosition unity.Vec3 // player's current position
}
