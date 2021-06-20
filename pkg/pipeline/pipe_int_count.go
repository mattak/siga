package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
)

type CountPipeInt struct {
}

func (c CobraCommandInput) CreateCountPipeInt() CountPipeInt {
	return CountPipeInt{}
}

func (pipe CountPipeInt) Execute(df *dataframe.DataFrame) int {
	return len(df.Labels)
}
