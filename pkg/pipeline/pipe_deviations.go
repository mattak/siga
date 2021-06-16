package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type DeviationsCommandOption struct {
	ColumnName string
	Spans      []int
	Output     OutputOption
}

type DeviationsCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    DeviationsCommandOption
}

func (c CobraCommandInput) CreateDeviationsCommandOption(option OutputOption) DeviationsCommandOption {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}
	columnName := c.Args[0]

	spans := make([]int, len(c.Args)-1)
	for i := 0; i < len(c.Args)-1; i++ {
		spans[i] = util.ParseInt(c.Args[i+1])
	}

	return DeviationsCommandOption{
		ColumnName: columnName,
		Spans:      spans,
		Output:     option,
	}
}

func (option DeviationsCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return DeviationsCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe DeviationsCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	vector, err := df.ExtractColumn(pipe.Option.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	vector.Reverse()

	for i := 0; i < len(pipe.Option.Spans); i++ {
		span := pipe.Option.Spans[i]
		if span <= 0 {
			log.Fatalf("SPAN should be more than 1: %d\n", span)
		}

		line := vector.Deviations(span)
		line.Reverse()

		label := pipe.Option.Output.ColumnName
		if label == "" {
			label = fmt.Sprintf("deviation_%s_%d", pipe.Option.ColumnName, span)
		}

		err = df.AddColumn(label, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	return df
}
