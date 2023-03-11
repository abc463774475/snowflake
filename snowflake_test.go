package snowflake

import "testing"

func TestGetID(t *testing.T) {
	Init(1023)

	for i := 0; i < 100; i++ {
		t.Log(GetID())
	}

	t.Log("done")
}
