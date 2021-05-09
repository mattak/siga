package main

import (
	"github.com/mattak/siga/cmd"
	"testing"
)

func TestSelectRecords(t *testing.T) {
	df := &cmd.DataFrame{}
	df.Headers = []string{"date", "column1"}
	df.Labels = []string{"seq1", "seq2", "seq3", "seq4"}
	df.Data = [][]float64{
		{1, 1},
		{1, 0},
		{2, 2},
		{2, 0},
	}

	df.SelectRecords(0, 2)

	if len(df.Labels) != 2 {
		t.Fatal("labels length is not expected")
	}
	if len(df.Data) != 2 {
		t.Fatal("data length is not expected")
	}
	if df.Data[0][0] != 1.0 || df.Data[0][1] != 1.0 {
		t.Fatal("data[0][0],data[0][1] is not expected")
	}
	if df.Data[1][0] != 2.0 || df.Data[1][1] != 2.0 {
		t.Fatalf("data[1][0],data[1][1] is not expected: %f, %f", df.Data[1][0], df.Data[1][1])
	}
}
