package main

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"math"
	"testing"
)

func expectInt(t *testing.T, name string, real int, expect int) {
	if real != expect {
		t.Fatalf("%s is not matched: %d <=> %d", name, real, expect)
	}
}

func expectValue(t *testing.T, name string, real float64, expect float64) {
	if math.IsNaN(expect) && math.IsNaN(real) || math.Abs(real - expect) <= 1e-9 {
		return
	}
	t.Fatalf("%s is not matched: %f <=> %f", name, real, expect)
}

func expectAverageLoss(t *testing.T, name string, real toolkit.AverageVariance, expect toolkit.AverageVariance) {
	expectValue(t, fmt.Sprintf("%s.Total", name), expect.Total, real.Total)
	if expect.Count != real.Count {
		t.Fatalf("%s.Count is not matched: %d <=> %d", name, expect.Count, real.Count)
	}
	expectValue(t, fmt.Sprintf("%s.Average", name), expect.Average, real.Average)
	expectValue(t, fmt.Sprintf("%s.Variance", name), expect.Variance, real.Variance)
}

