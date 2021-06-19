package pipeline

import (
	"fmt"
	"github.com/mattak/siga/pkg/dataframe"
	"os"
)

type LogCommandOption struct {
}

type LogCommandPipe struct {
	DataFrame *dataframe.DataFrame
}

func (c CobraCommandInput) CreateLogCommandOption(option OutputOption) LogCommandOption {
	return LogCommandOption{}
}

func (option LogCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return LogCommandPipe{
		DataFrame: df,
	}
}

func (pipe LogCommandPipe) Execute() *dataframe.DataFrame {
	df := pipe.DataFrame
	_, _ = fmt.Fprintln(os.Stderr, df.ToTsvString(false))
	return df
}
