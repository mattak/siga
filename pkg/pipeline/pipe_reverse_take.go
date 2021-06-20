package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type ReverseTakePipe struct {
	Size int
}

func (c CobraCommandInput) CreateReverseTakePipe(option OutputOption) ReverseTakePipe {
	if len(c.Args) < 1 {
		log.Fatal("SIZE should be declared")
	}

	size := util.ParseInt(c.Args[0])

	return ReverseTakePipe{
		Size: size,
	}
}

func (pipe ReverseTakePipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	df.Reverse()
	df.Take(pipe.Size)
	df.Reverse()
	return df
}
