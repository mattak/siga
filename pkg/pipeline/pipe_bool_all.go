package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type DetectAllPipeBool struct {
	ColumnName  string
	DetectValue float64
}

func (c CobraCommandInput) CreateDetectAllPipeBool() DetectAllPipeBool {
	if len(c.Args) < 1 {
		log.Fatal("COLUMN_NAME must be declared")
	}
	columnName := c.Args[0]
	detectValue := 1.0
	if len(c.Args) >= 2 {
		detectValue = util.ParseFloat64(c.Args[1])
	}
	return DetectAllPipeBool{
		ColumnName:  columnName,
		DetectValue: detectValue,
	}
}

func (pipe DetectAllPipeBool) Execute(df *dataframe.DataFrame) bool {
	vector, err := df.ExtractColumn(pipe.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	result := vector.IsAllValue(pipe.DetectValue)
	return result
}
