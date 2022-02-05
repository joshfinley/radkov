package unity_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

func TestNewUnityGame(t *testing.T) {
	offsetGameObjectManager := uintptr(0x17F8D28)
	ug, err := unity.NewUnityGame("EscapeFromTarkov.exe", offsetGameObjectManager)
	if err != nil {
		t.Fail()
	}

	if ug == nil {
		t.FailNow()
	}
}
