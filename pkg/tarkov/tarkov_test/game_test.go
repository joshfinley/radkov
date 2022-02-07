package tarkov_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/tarkov"
)

func TestGameMain(t *testing.T) {
	tg, err := tarkov.NewTarkovGame()
	if err != nil {
		t.FailNow()
	}

	if err != nil {
		t.FailNow()
	}

	t.Log(tg)
}
