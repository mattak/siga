package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type AddCommandOption struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

type AddCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    AddCommandOption
}

func (c CobraCommandInput) CreateAddCommandOption(option OutputOption) AddCommandOption {
	if len(c.Args) < 1 {
		log.Fatal("More than 1 COLUMN_NAME should be declared")
	}

	columnNameOrValues := c.Args

	return AddCommandOption{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (option AddCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return AddCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe AddCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.Option.ColumnNameOrValues)
	vector := matrix.Add()
	label := pipe.Option.Output.ColumnName

	if label == "" {
		label = fmt.Sprintf("add_%s", strings.Join(pipe.Option.ColumnNameOrValues, "_"))
	}

	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
