package stat

import "testing"

func TestDistribution(t *testing.T) {

	f := Float64Slice([]float64{1, 2, 3})

	if x := Freedom(f, false); x != 2 {
		t.Errorf("Degrees of freedom for one-tail is invalid = %f, want %d", x, 2)
	}
	if x := Freedom(f, true); x != 1 {
		t.Errorf("Degrees of freedom for two-tail is invalid = %f, want %d", x, 1)
	}

}
