package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type AnomalyCommandInput struct {
	DataFrame *dataframe.DataFrame
	Option    AnomalyCommandOption
}

type AnomalyCommandOption struct {
	ColumnName string
	Span       int
	Output     OutputOption
}

type OutputOption struct {
	ColumnName string
}

func (c CobraCommandInput) CreateAnomalyCommandOption(option OutputOption) AnomalyCommandOption {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnName := c.Args[0]
	span := util.ParseInt(c.Args[1])

	return AnomalyCommandOption{
		ColumnName: columnName,
		Span:       span,
		Output:     option,
	}
}

func (option AnomalyCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return AnomalyCommandInput{
		Option:    option,
		DataFrame: df,
	}
}

func (input AnomalyCommandInput) Execute() *dataframe.DataFrame {
	columnName := input.Option.ColumnName
	span := input.Option.Span
	outputColumnName := input.Option.ColumnName
	df := input.DataFrame

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
