package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestVectorReverse(t *testing.T) {
	data := cmd.Vector{1,2,3}
	data.Reverse()

	if data[0] != 3 || data[1] != 2 || data[2] != 1 {
		t.Fatal("data is not reversed")
	}
}
