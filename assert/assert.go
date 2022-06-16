package assert

import "testing"

func Equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %[1]T %[1]v to equal %[2]T %[2]v", a, b)
	}

}
