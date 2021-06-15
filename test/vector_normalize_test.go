package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"testing"
)

func TestNormalize(t *testing.T) {
	t.Run("same", func(t *testing.T) {
		data := dataframe.Vector{1, 2, 3}
		normalize := data.NormalizeByStart()
		ExpectInt(t, "len", len(normalize), 3)
		ExpectValue(t, "normalize[0]", normalize[0], 1)
		ExpectValue(t, "normalize[1]", normalize[1], 2)
		ExpectValue(t, "normalize[2]", normalize[2], 3)
	})
	t.Run("empty", func(t *testing.T) {
		data := dataframe.Vector{}
		normalize := data.NormalizeByStart()
		ExpectInt(t, "len", len(normalize), 0)
	})
	t.Run("half", func(t *testing.T) {
		data := dataframe.Vector{4, 2, 1}
		normalize := data.NormalizeByStart()
		ExpectInt(t, "len", len(normalize), 3)
		ExpectValue(t, "normalize[0]", normalize[0], 1)
		ExpectValue(t, "normalize[1]", normalize[1], 0.5)
		ExpectValue(t, "normalize[2]", normalize[2], 0.25)
	})
}
