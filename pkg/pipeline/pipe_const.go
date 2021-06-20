package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type ConstPipe struct {
	Values []float64
	Output OutputOption
}

func (c CobraCommandInput) CreateConstPipe(option OutputOption) ConstPipe {
	if len(c.Args) < 1 {
		log.Fatal("NUMBER should be declared")
	}

	values := make([]float64, len(c.Args))
	for i := 0; i < len(values); i++ {
		values[i] = util.ParseFloat64(c.Args[i])
	}

	return ConstPipe{
		Values: values,
		Output: option,
	}
}

func (pipe ConstPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	for i := 0; i < len(pipe.Values); i++ {
		n := pipe.Values[i]
		vector := dataframe.CreateVector(len(df.Labels))
		vector.Fill(n)

		label := pipe.Output.ColumnName
		if label == "" {
			label = fmt.Sprintf("const_%.3f", n)
		}

		err := df.AddColumn(label, vector)
		if err != nil {
			log.Fatal(err)
		}
	}

	return df
}
