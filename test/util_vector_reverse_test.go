package main

import (
	"github.com/mattak/siga/pkg"
	"testing"
)

func TestVectorReverse(t *testing.T) {
	data := pkg.Vector{1,2,3}
	data.Reverse()

	if data[0] != 3 || data[1] != 2 || data[2] != 1 {
		t.Fatal("data is not reversed")
	}
}
