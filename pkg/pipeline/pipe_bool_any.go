package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type DetectAnyPipeBool struct {
	ColumnName  string
	DetectValue float64
}

func (c CobraCommandInput) CreateDetectAnyPipeBool() DetectAnyPipeBool {
	if len(c.Args) < 1 {
		log.Fatal("COLUMN_NAME must be declared")
	}
	columnName := c.Args[0]
	detectValue := 1.0
	if len(c.Args) >= 2 {
		detectValue = util.ParseFloat64(c.Args[1])
	}
	return DetectAnyPipeBool{
		ColumnName:  columnName,
		DetectValue: detectValue,
	}
}

func (pipe DetectAnyPipeBool) Execute(df *dataframe.DataFrame) bool {
	vector, err := df.ExtractColumn(pipe.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	result := vector.HasAnyValue(pipe.DetectValue)
	return result
}
