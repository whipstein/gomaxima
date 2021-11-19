package maxima

import "testing"

func TestBuildMatrix(t *testing.T) {
	outvals := []string{"x", "y", "z", "a", "b", "c"}

	tgt := "matrix([x,y,z],[a,b,c])"
	out := BuildMatrix(2, 3, outvals...)
	if out != tgt {
		t.Errorf("BuildMatrix failed: got %s want %s\n", out, tgt)
	}

	tgt = "matrix([x,y],[z,a],[b,c])"
	out = BuildMatrix(3, 2, outvals...)
	if out != tgt {
		t.Errorf("BuildMatrix failed: got %s want %s\n", out, tgt)
	}
}
