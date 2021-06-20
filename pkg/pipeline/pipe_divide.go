package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type DividePipe struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

func (c CobraCommandInput) CreateDividePipe(option OutputOption) DividePipe {
	if len(c.Args) < 1 {
		log.Fatal("More than 1 COLUMN_NAME should be declared")
	}

	columnNameOrValues := c.Args

	return DividePipe{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (pipe DividePipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.ColumnNameOrValues)
	vector := matrix.Divide()
	label := pipe.Output.ColumnName

	if label == "" {
		label = fmt.Sprintf("divide_%s", strings.Join(pipe.ColumnNameOrValues, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
