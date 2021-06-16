package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"log"
	"strings"
)

type LessEqualCommandOption struct {
	ColumnNameOrValues []string
	Output             OutputOption
}

type LessEqualCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    LessEqualCommandOption
}

func (c CobraCommandInput) CreateLessEqualCommandOption(option OutputOption) LessEqualCommandOption {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME, SPAN should be declared")
	}

	columnNameOrValues := c.Args

	return LessEqualCommandOption{
		ColumnNameOrValues: columnNameOrValues,
		Output:             option,
	}
}

func (option LessEqualCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return LessEqualCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe LessEqualCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	matrix := df.ExtractMatrixByColumnNameOrValue(pipe.Option.ColumnNameOrValues)
	label := pipe.Option.Output.ColumnName

	vector := matrix.LessEqual()
	if label == "" {
		label = fmt.Sprintf("le_%s", strings.Join(pipe.Option.ColumnNameOrValues, "_"))
	}
	err := df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}

	return df
}
