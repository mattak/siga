package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type LessEqualPipe struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

func (c CobraCommandInput) CreateLessEqualPipe(option OutputOption) LessEqualPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return LessEqualPipe{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (pipe LessEqualPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.ColumnNameOrValues)
	label := pipe.Output.ColumnName

	vector := matrix.LessEqual()
	if label == "" {
		label = fmt.Sprintf("le_%s", strings.Join(pipe.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
