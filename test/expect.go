package main

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func ExpectInt(t *testing.T, name string, real int, expect int) {
	if real != expect {
		t.Fatalf("%s is not matched: %d <=> %d", name, real, expect)
	}
}

func ExpectString(t *testing.T, name string, real string, expect string) {
	if real != expect {
		t.Fatalf("%s is not matched: %s <=> %s", name, real, expect)
	}
}

func ExpectValue(t *testing.T, name string, real float64, expect float64) {
	if math.IsNaN(expect) && math.IsNaN(real) || math.Abs(real-expect) <= 1e-9 {
		return
	}
	t.Fatalf("%s is not matched: %f <=> %f", name, real, expect)
}

func ExpectAverageLoss(t *testing.T, name string, real dataframe.AverageVariance, expect dataframe.AverageVariance) {
	ExpectValue(t, fmt.Sprintf("%s.Total", name), expect.Total, real.Total)
	if expect.Count != real.Count {
		t.Fatalf("%s.Count is not matched: %d <=> %d", name, expect.Count, real.Count)
	}
	ExpectValue(t, fmt.Sprintf("%s.Average", name), expect.Average, real.Average)
	ExpectValue(t, fmt.Sprintf("%s.Variance", name), expect.Variance, real.Variance)
}

func ExpectRangeValue(t *testing.T, name string, real float64, expectFrom float64, expectTo float64) {
	if real < expectFrom || real > expectTo {
		t.Fatal("wrong range of value ", name)
	}
}
