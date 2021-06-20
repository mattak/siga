package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"os"
)

type LogCommandPipe struct {
}

func (c CobraCommandInput) CreateLogCommandPipe(option OutputOption) LogCommandPipe {
	return LogCommandPipe{}
}

func (optoin LogCommandPipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	_, _ = fmt.Fprintln(os.Stderr, df.ToTsvString(false))
	return df
}
