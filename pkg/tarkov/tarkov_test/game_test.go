package tarkov_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
)

func TestGameMain(t *testing.T) {
	tg, err := tarkov.LoadGame()
	if err != nil {
		t.FailNow()
	}

	err = tg.GameMain()
	if err != nil {
		t.FailNow()
	}

	t.Log(tg)
}
