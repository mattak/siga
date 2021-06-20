package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type ReverseTakeCommandOption struct {
	Size int
}

type ReverseTakeCommandPipe struct {
	DataFrame *dataframe.DataFrame
	Option    ReverseTakeCommandOption
}

func (c CobraCommandInput) CreateReverseTakeCommandOption(option OutputOption) ReverseTakeCommandOption {
	if len(c.Args) < 1 {
		log.Fatal("SIZE should be declared")
	}

	size := util.ParseInt(c.Args[0])

	return ReverseTakeCommandOption{
		Size: size,
	}
}

func (option ReverseTakeCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return ReverseTakeCommandPipe{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe ReverseTakeCommandPipe) Execute() *dataframe.DataFrame {
	pipe.DataFrame.Reverse()
	pipe.DataFrame.Take(pipe.Option.Size)
	pipe.DataFrame.Reverse()
	return pipe.DataFrame
}
