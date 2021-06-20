package main

import (
	"github.com/mattak/siga/pkg/dataframe"
	"testing"
)

func TestCreateDataFrame(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		df, err := dataframe.CreateDataFrame([]string{"value"}, []dataframe.Vector{
			[]float64{1, 2, 3},
		})
		if err != nil {
			t.Fatal("error should not be raised")
		}
		ExpectInt(t, "labels length", len(df.Labels), 3)
		ExpectString(t, "label[0]", df.Labels[0], "1")
		ExpectString(t, "label[1]", df.Labels[1], "2")
		ExpectString(t, "label[1]", df.Labels[2], "3")
		ExpectInt(t, "headers length", len(df.Headers), 2)
		ExpectString(t, "header[0]", df.Headers[0], "index")
		ExpectString(t, "header[1]", df.Headers[1], "value")
		ExpectInt(t, "data length", len(df.Data), 3)
		ExpectInt(t, "data[0] length", len(df.Data[0]), 1)
		ExpectValue(t, "data[0][0]", df.Data[0][0], 1.0)
		ExpectValue(t, "data[0][1]", df.Data[1][0], 2.0)
		ExpectValue(t, "data[0][2]", df.Data[2][0], 3.0)
	})
	t.Run("header data is not matched", func(t *testing.T) {
		df, err := dataframe.CreateDataFrame([]string{"value1", "value2"}, []dataframe.Vector{
			[]float64{1, 2, 3},
		})
		if err == nil {
			t.Fatal("error must be raised")
		}
		if df != nil {
			t.Fatal("dataframe must be nil")
		}
	})
}

func TestCreateDataFrameWithLabels(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		df, err := dataframe.CreateDataFrameWithLabels(
			[]string{"weekday", "value"},
			[]string{"Mon", "Tue", "Wed"},
			[]dataframe.Vector{
				[]float64{1, 2, 3},
			})
		if err != nil {
			t.Fatalf("error should no be raised: %v\n", err)
		}
		ExpectInt(t, "labels length", len(df.Labels), 3)
		ExpectString(t, "label[0]", df.Labels[0], "Mon")
		ExpectString(t, "label[1]", df.Labels[1], "Tue")
		ExpectString(t, "label[1]", df.Labels[2], "Wed")
		ExpectInt(t, "headers length", len(df.Headers), 2)
		ExpectString(t, "header[0]", df.Headers[0], "weekday")
		ExpectString(t, "header[1]", df.Headers[1], "value")
		ExpectInt(t, "data length", len(df.Data), 3)
		ExpectInt(t, "data[0] length", len(df.Data[0]), 1)
		ExpectValue(t, "data[0][0]", df.Data[0][0], 1.0)
		ExpectValue(t, "data[0][1]", df.Data[1][0], 2.0)
		ExpectValue(t, "data[0][2]", df.Data[2][0], 3.0)
	})
	t.Run("header and data length not matched", func(t *testing.T) {
		df, err := dataframe.CreateDataFrameWithLabels(
			[]string{"weekday"},
			[]string{"Mon", "Tue", "Wed"},
			[]dataframe.Vector{
				[]float64{1, 2, 3},
			})
		if err == nil {
			t.Fatal("error should no be raised")
		}
		if df != nil {
			t.Fatal("DataFrame should be nil")
		}
	})
	t.Run("labels and data length not matched", func(t *testing.T) {
		df, err := dataframe.CreateDataFrameWithLabels(
			[]string{"weekday", "value"},
			[]string{"Mon", "Tue"},
			[]dataframe.Vector{
				[]float64{1, 2, 3},
			})
		if err == nil {
			t.Fatal("error should no be raised")
		}
		if df != nil {
			t.Fatal("DataFrame should be nil")
		}
	})
}

func TestClone(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		df := dataframe.DataFrame{
			[]string{},
			[]string{},
			[][]float64{},
		}
		df2 := df.Clone()
		ExpectInt(t, "labels length", len(df2.Labels), 0)
		ExpectInt(t, "headers length", len(df2.Headers), 0)
		ExpectInt(t, "data length", len(df2.Data), 0)
	})
	t.Run("clone", func(t *testing.T) {
		df := dataframe.DataFrame{
			[]string{"index", "value"},
			[]string{"1", "2"},
			[][]float64{
				[]float64{10},
				[]float64{20},
			},
		}
		df2 := df.Clone()
		ExpectInt(t, "labels length", len(df2.Labels), 2)
		ExpectInt(t, "headers length", len(df2.Headers), 2)
		ExpectInt(t, "data length", len(df2.Data), 2)
		ExpectInt(t, "data[0] length", len(df2.Data[0]), 1)
		ExpectInt(t, "data[1] length", len(df2.Data[1]), 1)
		ExpectValue(t, "data[0][0]", df2.Data[0][0], 10.0)
		ExpectValue(t, "data[1][0]", df2.Data[1][0], 20.0)
	})
}
