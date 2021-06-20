package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type GreaterThanPipe struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

func (c CobraCommandInput) CreateGreaterThanPipe(option OutputOption) GreaterThanPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return GreaterThanPipe{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (pipe GreaterThanPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.ColumnNameOrValues)
	label := pipe.Output.ColumnName

	vector := matrix.GreaterThan()
	if label == "" {
		label = fmt.Sprintf("gt_%s", strings.Join(pipe.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
