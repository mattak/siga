package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type LessThanPipe struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

func (c CobraCommandInput) CreateLessThanPipe(option OutputOption) LessThanPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return LessThanPipe{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (pipe LessThanPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.ColumnNameOrValues)
	label := pipe.Output.ColumnName

	vector := matrix.LessThan()
	if label == "" {
		label = fmt.Sprintf("lt_%s", strings.Join(pipe.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
