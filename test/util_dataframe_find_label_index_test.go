package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestFindLabelIndex(t *testing.T) {
	df := &cmd.DataFrame{}
	df.Headers = []string{"date", "column1", "column2"}
	df.Labels = []string{"seq1", "seq2", "seq3"}
	df.Data = [][]float64{
		{1, 10},
		{2, 20},
		{3, 30},
	}

	if df.FindLabelIndex("seq1") != 0 {
		t.Fatal("search result is not expected: seq1")
	}
	if df.FindLabelIndex("seq3") != 2 {
		t.Fatal("search result is not expected: seq3")
	}
	if df.FindLabelIndex("aaa") != -1 {
		t.Fatal("search result is not expected")
	}
}
