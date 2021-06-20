package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"log"
)

type TakePipe struct {
	Size int
}

func (c CobraCommandInput) CreateTakePipe(option OutputOption) TakePipe {
	if len(c.Args) < 1 {
		log.Fatal("SIZE should be declared")
	}

	size := util.ParseInt(c.Args[0])

	return TakePipe{
		Size: size,
	}
}

func (pipe TakePipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	df.Take(pipe.Size)
	return df
}
