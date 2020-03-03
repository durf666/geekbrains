package statistic

import "testing"

type testpair struct {
	values  []float64
	average float64
}
type testpair2 struct {
	values []float64
	sum    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var sum_tests = []testpair2{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
func TestSumeSet(t *testing.T) {
	for _, pair := range sum_tests {
		v := Sum(pair.values)
		if v != pair.sum {
			t.Error(
				"For", pair.values,
				"expected", pair.sum,
				"got", v,
			)
		}
	}
}
