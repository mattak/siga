package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type GreaterEqualCommandOption struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

type GreaterEqualCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    GreaterEqualCommandOption
}

func (c CobraCommandInput) CreateGreaterEqualCommandOption(option OutputOption) GreaterEqualCommandOption {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return GreaterEqualCommandOption{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (option GreaterEqualCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return GreaterEqualCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe GreaterEqualCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.Option.ColumnNameOrValues)
	label := pipe.Option.Output.ColumnName

	vector := matrix.GreaterEqual()
	if label == "" {
		label = fmt.Sprintf("ge_%s", strings.Join(pipe.Option.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
