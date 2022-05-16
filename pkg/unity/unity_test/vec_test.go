package unity_test

import (
	"testing"

	"github.com/joshfinley/radkov/pkg/unity"
)

func TestMarshalVec3(t *testing.T) {
	v := unity.Vec3{
		X: 1,
		Y: 2,
		Z: 3,
	}

	m := unity.MarshalVec3(v)
	u := unity.UnmarshalVec3(m)

	if u != v {
		t.FailNow()
	}
}
