package stat

import (
	"math"
	"testing"
)

func TestTtest(t *testing.T) {
	f1, f2 := setup_ttest()
	f, ff := Float64Slice(f1), Float64Slice(f2)

	if tval := TTest(f, ff); tval-10.2321 > math.SmallestNonzeroFloat64 {
		t.Errorf("Invalid tvalue has: %f, wanted %f", tval, 10.232033)
	}

	if pvalue := TTestpvalue(f, ff, false); pvalue < math.SmallestNonzeroFloat64 {
		t.Errorf("Invalid pvalue has: %f, wanted %f", pvalue, 0.0)
	}

	f, ff = Float64Slice([]float64{10, 20, 30, 40, 50, 60}),
		Float64Slice([]float64{20, 30, 40, 50, 60, 70})

	if pvalue := TTestpvalue(f, ff, false); pvalue-0.18817 > math.SmallestNonzeroFloat64 {
		t.Errorf("Invalid pvalue has: %f, wanted %f", pvalue, 0.18817)
	}

	//Wikipedia example
	f, ff =
		Float64Slice([]float64{30.02, 29.99, 30.11, 29.97, 30.01, 29.99}),
		Float64Slice([]float64{29.89, 29.93, 29.72, 29.98, 30.02, 29.98})

	if tval := TTest(f, ff); tval-1.959006 > math.SmallestNonzeroFloat64 {
		t.Errorf("Invalid tvalue has: %f, wanted %f", tval, 1.959006)
	}

	if pvalue := TTestpvalue(f, ff, false); pvalue != 0.045 {
		t.Errorf("Invalid pvalue has: %f, wanted %f", pvalue, 0.045)
	}
}

func setup_ttest() ([]float64, []float64) {
	var f1 = []float64{
		246915,
		263093,
		210897,
		263092,
		260268,
		263095,
		263095,
		263094,
		263094,
		263095,
		263095,
		263094,
		263094,
		263094,
		263095,
		263094,
		263095,
		263093,
		263095,
		263094,
		258373,
		257810,
		263095,
		263095,
		263094,
		263095,
		261008,
		263095,
		261387,
		263079,
		263095,
		263094,
		263094,
		251209,
		263095,
		259320,
		263095,
		263095,
		263094,
		263095,
		263095,
		263095,
		263094,
		263093,
		263093,
		263092,
		263095,
		255271,
		263094,
		263095,
		263095,
		259322,
		263094,
		261041,
		263095,
		263095,
		263095,
		263095,
		262693,
		263093,
	}

	var f2 = []float64{
		240798,
		240796,
		240797,
		240495,
		240796,
		240797,
		240796,
		240796,
		240796,
		239099,
		240797,
		240794,
		240795,
		240797,
		240797,
		240796,
		237023,
		240797,
		240797,
		240797,
		114693,
		239100,
		240781,
		240796,
		240797,
		232943,
		237504,
		240796,
		235727,
		237971,
		240796,
		240795,
		240797,
		240797,
		240797,
		240797,
		240797,
		240797,
		240797,
		240796,
		236567,
		224618,
		240796,
		240796,
		210283,
		240796,
		240797,
		236425,
		240795,
		240796,
		224618,
		240797,
		240796,
		240796,
		240795,
		240796,
		238385,
		240796,
		240797,
		240796,
	}
	return f1, f2
}
