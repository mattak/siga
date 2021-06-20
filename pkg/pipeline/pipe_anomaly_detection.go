package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type AnomalyPipe struct {
	ColumnName string
	Span       int
	Output     OutputOption
}

func (c CobraCommandInput) CreateAnomalyPipe(option OutputOption) AnomalyPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnName := c.Args[0]
	span := util.ParseInt(c.Args[1])

	return AnomalyPipe{
		ColumnName: columnName,
		Span:       span,
		Output:     option,
	}
}

func (pipe AnomalyPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	columnName := pipe.ColumnName
	span := pipe.Span
	outputColumnName := pipe.Output.ColumnName

	vector, err := df.ExtractColumn(columnName)
	if err != nil {
		log.Fatalf("Not found index of header name: %s\n", columnName)
	}
	vector.Reverse()

	result := vector.SigmaAnomalies(span)
	result.Reverse()

	if outputColumnName == "" {
		outputColumnName = fmt.Sprintf("%s_anomaly", columnName)
	}

	err = df.AddColumn(outputColumnName, result)
	if err != nil {
		log.Fatalf("add result column failed\n")
	}

	return df
}
