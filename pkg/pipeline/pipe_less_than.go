package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type LessThanCommandOption struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

type LessThanCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    LessThanCommandOption
}

func (c CobraCommandInput) CreateLessThanCommandOption(option OutputOption) LessThanCommandOption {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return LessThanCommandOption{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (option LessThanCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return LessThanCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe LessThanCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.Option.ColumnNameOrValues)
	label := pipe.Option.Output.ColumnName

	vector := matrix.LessThan()
	if label == "" {
		label = fmt.Sprintf("lt_%s", strings.Join(pipe.Option.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
