package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type DetectAnyCommandOption struct {
	ColumnName  string
	DetectValue float64
}

type DetectAnyCommandPipeOutput struct {
	DataFrame *dataframe.DataFrame
	Option    DetectAnyCommandOption
}

func (c CobraCommandInput) CreateDetectAnyCommandOption() DetectAnyCommandOption {
	if len(c.Args) < 1 {
		log.Fatal("COLUMN_NAME must be declared")
	}
	columnName := c.Args[0]
	detectValue := 1.0
	if len(c.Args) >= 2 {
		detectValue = util.ParseFloat64(c.Args[1])
	}
	return DetectAnyCommandOption{
		ColumnName:  columnName,
		DetectValue: detectValue,
	}
}

func (option DetectAnyCommandOption) CreatePipeBool(df *dataframe.DataFrame) PipeBool {
	return DetectAnyCommandPipeOutput{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe DetectAnyCommandPipeOutput) Execute() bool {
	vector, err := pipe.DataFrame.ExtractColumn(pipe.Option.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	result := vector.HasAnyValue(pipe.Option.DetectValue)
	return result
}
