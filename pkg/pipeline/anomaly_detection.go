package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type AnomalyCommandInput struct {
	ColumnName string
	Span       int
	DataFrame  dataframe.DataFrame
	Option     AnomalyCommandOption
}

type AnomalyCommandOption struct {
	ColumnName string
}

func (c CobraCommandInput) CreateAnomalyCommandInput(option AnomalyCommandOption) AnomalyCommandInput {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnName := c.Args[0]
	span := util.ParseInt(c.Args[1])
	df := dataframe.ReadDataFrameByStdinTsv()

	return AnomalyCommandInput{
		ColumnName: columnName,
		Span:       span,
		DataFrame:  *df,
		Option:     option,
	}
}

func (input AnomalyCommandInput) Execute() *dataframe.DataFrame {
	columnName := input.ColumnName
	span := input.Span
	df := input.DataFrame
	outputColumnName := input.Option.ColumnName

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

	return &df
}
