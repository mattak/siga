package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"math"
)

type SelectLastValuePipeFloat64 struct {
	ColumnName string
}

func (c CobraCommandInput) CreateSelectLastValuePipeFloat64() SelectLastValuePipeFloat64 {
	if len(c.Args) < 1 {
		log.Fatal("COLUMN_NAME required")
	}

	return SelectLastValuePipeFloat64{
		ColumnName: c.Args[0],
	}
}

func (pipe SelectLastValuePipeFloat64) Execute(df *dataframe.DataFrame) float64 {
	vector, err := df.ExtractColumn(pipe.ColumnName)
	if err != nil {
		log.Fatalf("ERROR: not found column %v\n", err)
	}
	if len(vector) < 1 {
		return math.NaN()
	}
	return vector[len(vector)-1]
}
