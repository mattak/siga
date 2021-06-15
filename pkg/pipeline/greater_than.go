package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type GreaterThanCommandOption struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

type GreaterThanCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    GreaterThanCommandOption
}

func (c CobraCommandInput) CreateGreaterThanCommandOption(option OutputOption) GreaterThanCommandOption {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return GreaterThanCommandOption{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (option GreaterThanCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return GreaterThanCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe GreaterThanCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.Option.ColumnNameOrValues)
	label := pipe.Option.Output.ColumnName

	vector := matrix.GreaterThan()
	if label == "" {
		label = fmt.Sprintf("gt_%s", strings.Join(pipe.Option.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
