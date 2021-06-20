package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type AddPipe struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

func (c CobraCommandInput) CreateAddPipe(option OutputOption) AddPipe {
	if len(c.Args) < 1 {
		log.Fatal("More than 1 COLUMN_NAME should be declared")
	}

	columnNameOrValues := c.Args

	return AddPipe{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (pipe AddPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.ColumnNameOrValues)
	vector := matrix.Add()
	label := pipe.Output.ColumnName

	if label == "" {
		label = fmt.Sprintf("add_%s", strings.Join(pipe.ColumnNameOrValues, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
