package main

import (
	"github.com/mattak/siga/toolkit"
	"math"
	"testing"
)

func TestVectorIsAllValue(t *testing.T) {
	data1 := toolkit.Vector{1, 1}
	if data1.IsAllValue(0) {
		t.Fatal("0 should not be detected")
	}
	if !data1.IsAllValue(1) {
		t.Fatal("1 should be detected")
	}
	data2 := toolkit.Vector{math.NaN(), math.NaN()}
	if data2.IsAllValue(1) {
		t.Fatal("1 should not be detected")
	}
	if !data2.IsAllValue(math.NaN()) {
		t.Fatal("NaN should be detected")
	}
	data3 := toolkit.Vector{1, 2}
	if data3.IsAllValue(1) {
		t.Fatal("1 should not be detected")
	}
}
