package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type ConstCommandOption struct {
	Values []float64
	Output OutputOption
}

type ConstCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    ConstCommandOption
}

func (c CobraCommandInput) CreateConstCommandOption(option OutputOption) ConstCommandOption {
	if len(c.Args) < 1 {
		log.Fatal("NUMBER should be declared")
	}

	values := make([]float64, len(c.Args))
	for i := 0; i < len(values); i++ {
		values[i] = util.ParseFloat64(c.Args[i])
	}

	return ConstCommandOption{
		Values: values,
		Output: option,
	}
}

func (option ConstCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return ConstCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe ConstCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame

	for i := 0; i < len(pipe.Option.Values); i++ {
		n := pipe.Option.Values[i]
		vector := dataframe.CreateVector(len(df.Labels))
		vector.Fill(n)

		label := pipe.Option.Output.ColumnName
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
