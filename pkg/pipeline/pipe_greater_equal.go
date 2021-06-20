package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type GreaterEqualPipe struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

func (c CobraCommandInput) CreateGreaterEqualPipe(option OutputOption) GreaterEqualPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return GreaterEqualPipe{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (pipe GreaterEqualPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.ColumnNameOrValues)
	label := pipe.Output.ColumnName

	vector := matrix.GreaterEqual()
	if label == "" {
		label = fmt.Sprintf("ge_%s", strings.Join(pipe.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
