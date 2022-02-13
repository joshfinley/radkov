package unity_test

import (
	"testing"

	"gitlab.clan-ac.xyz/ac-gameworx/radkov/pkg/unity"
)

func TestMarshalVec3(t *testing.T) {
	v := unity.Vec3{
		X: 1,
		Y: 2,
		Z: 3,
	}

	m := v.Marshal()
	u := m.Unmarshal()

	if u != v {
		t.FailNow()
	}
}
