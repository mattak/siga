package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type DetectAllCommandOption struct {
	ColumnName  string
	DetectValue float64
}

type DetectAllCommandPipeOutput struct {
	DataFrame *dataframe.DataFrame
	Option    DetectAllCommandOption
}

func (c CobraCommandInput) CreateDetectAllCommandOption() DetectAllCommandOption {
	if len(c.Args) < 1 {
		log.Fatal("COLUMN_NAME must be declared")
	}
	columnName := c.Args[0]
	detectValue := 1.0
	if len(c.Args) >= 2 {
		detectValue = util.ParseFloat64(c.Args[1])
	}
	return DetectAllCommandOption{
		ColumnName:  columnName,
		DetectValue: detectValue,
	}
}

func (option DetectAllCommandOption) CreatePipeBool(df *dataframe.DataFrame) PipeBool {
	return DetectAllCommandPipeOutput{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe DetectAllCommandPipeOutput) Execute() bool {
	vector, err := pipe.DataFrame.ExtractColumn(pipe.Option.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	result := vector.IsAllValue(pipe.Option.DetectValue)
	return result
}
