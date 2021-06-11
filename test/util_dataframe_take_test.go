package main

import (
	"github.com/mattak/siga/toolkit"
	"testing"
)

func TestDataFrameTake(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		df := &toolkit.DataFrame{}
		df.Headers = []string{"date", "column1", "column2"}
		df.Labels = []string{"seq1", "seq2"}
		df.Data = [][]float64{
			{1, 10},
			{2, 20},
		}

		df.Take(10)
		if len(df.Labels) != 2 || len(df.Data) != 2 {
			t.Fatal("date size is not expected")
		}
		if df.Labels[0] != "seq1" || df.Labels[1] != "seq2" {
			t.Fatal("labels is not expected")
		}
		if df.Data[0][0] != 1 || df.Data[0][1] != 10 {
			t.Fatal("data[0] is not expected")
		}
		if df.Data[1][0] != 2 || df.Data[1][1] != 20 {
			t.Fatal("data[1] is not expected")
		}
	})

	t.Run("just", func(t *testing.T) {
		df := &toolkit.DataFrame{}
		df.Headers = []string{"date", "column1", "column2"}
		df.Labels = []string{"seq1", "seq2"}
		df.Data = [][]float64{
			{1, 10},
			{2, 20},
		}

		df.Take(2)
		if len(df.Labels) != 2 || len(df.Data) != 2 {
			t.Fatal("date size is not expected")
		}
		if df.Labels[0] != "seq1" || df.Labels[1] != "seq2" {
			t.Fatal("labels is not expected")
		}
		if df.Data[0][0] != 1 || df.Data[0][1] != 10 {
			t.Fatal("data[0] is not expected")
		}
		if df.Data[1][0] != 2 || df.Data[1][1] != 20 {
			t.Fatal("data[1] is not expected")
		}
	})

	t.Run("below", func(t *testing.T) {
		df := &toolkit.DataFrame{}
		df.Headers = []string{"date", "column1", "column2"}
		df.Labels = []string{"seq1", "seq2"}
		df.Data = [][]float64{
			{1, 10},
			{2, 20},
		}

		df.Take(1)
		if len(df.Labels) != 1 || len(df.Data) != 1 {
			t.Fatal("date size is not expected")
		}
		if df.Labels[0] != "seq1" {
			t.Fatal("labels is not expected")
		}
		if df.Data[0][0] != 1 || df.Data[0][1] != 10 {
			t.Fatal("data[0] is not expected")
		}
	})

	t.Run("zero", func(t *testing.T) {
		df := &toolkit.DataFrame{}
		df.Headers = []string{"date", "column1", "column2"}
		df.Labels = []string{"seq1", "seq2"}
		df.Data = [][]float64{
			{1, 10},
			{2, 20},
		}

		df.Take(0)
		if len(df.Labels) != 0 || len(df.Data) != 0 {
			t.Fatal("date size is not expected")
		}
	})
}
