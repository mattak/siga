package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type DeviationsPipe struct {
	ColumnName string
	Spans      []int
	Output     OutputOption
}

func (c CobraCommandInput) CreateDeviationsPipe(option OutputOption) DeviationsPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}
	columnName := c.Args[0]

	spans := make([]int, len(c.Args)-1)
	for i := 0; i < len(c.Args)-1; i++ {
		spans[i] = util.ParseInt(c.Args[i+1])
	}

	return DeviationsPipe{
		ColumnName: columnName,
		Spans:      spans,
		Output:     option,
	}
}

func (pipe DeviationsPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	vector, err := df.ExtractColumn(pipe.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	vector.Reverse()

	for i := 0; i < len(pipe.Spans); i++ {
		span := pipe.Spans[i]
		if span <= 0 {
			log.Fatalf("SPAN should be more than 1: %d\n", span)
		}

		line := vector.Deviations(span)
		line.Reverse()

		label := pipe.Output.ColumnName
		if label == "" {
			label = fmt.Sprintf("deviation_%s_%d", pipe.ColumnName, span)
		}

		err = df.AddColumn(label, line)
		if err != nil {
			log.Fatal(err)
		}
	}

	return df
}
