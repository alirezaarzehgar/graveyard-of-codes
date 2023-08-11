package cover

import "testing"

func TestAdd(t *testing.T) {
	v, _ := add(2, 3)
	if v != 5 {
		t.Error("fuck off")
	}
}
