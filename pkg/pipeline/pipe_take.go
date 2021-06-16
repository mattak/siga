package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type TakeCommandOption struct {
	Size int
}

type TakeCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    TakeCommandOption
}

func (c CobraCommandInput) CreateTakeCommandOption(option OutputOption) TakeCommandOption {
	if len(c.Args) < 1 {
		log.Fatal("SIZE should be declared")
	}

	size := util.ParseInt(c.Args[0])

	return TakeCommandOption{
		Size: size,
	}
}

func (option TakeCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return TakeCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe TakeCommandPipe) Execute() *dataframe.DataFrame {
	pipe.DataFrame.Take(pipe.Option.Size)
	return pipe.DataFrame
}
