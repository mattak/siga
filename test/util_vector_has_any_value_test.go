package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"math"
	"testing"
)

func TestVectorHasAnyValue(t *testing.T) {
	data := dataframe.Vector{1, 2, 3, math.NaN()}
	if data.HasAnyValue(0) {
		t.Fatal("0 should not be detected")
	}
	if !data.HasAnyValue(2) {
		t.Fatal("2 should be detected")
	}
	if !data.HasAnyValue(math.NaN()) {
		t.Fatal("NaN should be detected")
	}
}
