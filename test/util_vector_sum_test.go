package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestVectorSum(t *testing.T) {
	data := cmd.Vector{1,2,3}
	if data.Sum(0, 3) != 6 {
		t.Fatal("sum value is wrong")
	}
	if data.Sum(1, 2) != 5 {
		t.Fatal("sum value is wrong")
	}
	if data.Sum(2, 1) != 3 {
		t.Fatal("sum value is wrong")
	}
}
