package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type ShiftPipe struct {
	ColumnName string
	Offset     int
	Output     OutputOption
}

func (c CobraCommandInput) CreateShiftPipe(option OutputOption) ShiftPipe {
	if len(c.Args) < 2 {
		log.Fatal("COLUMN_NAME and OFFSET should be declared")
	}

	columnName := c.Args[0]
	offset := util.ParseInt(c.Args[1])

	return ShiftPipe{
		ColumnName: columnName,
		Offset:     offset,
		Output:     option,
	}
}

func (pipe ShiftPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	vector, err := df.ExtractColumn(pipe.ColumnName)
	if err != nil {
		log.Fatal(err)
	}
	vector = vector.Shift(pipe.Offset)

	label := pipe.Output.ColumnName
	if label == "" {
		label = fmt.Sprintf("shift_%s_%d", pipe.ColumnName, pipe.Offset)
	}
	err = df.AddColumn(label, vector)
	if err != nil {
		log.Fatal(err)
	}
	return df
}
