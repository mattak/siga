package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type Pipe interface {
	Execute(df *dataframe.DataFrame) *dataframe.DataFrame
}

type PipeCreator interface {
	CreatePipe() Pipe
}

type Pipeline []Pipe

func (creators Pipeline) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	for i := 0; i < len(creators); i++ {
		df = creators[i].Execute(df)
	}
	return df
}
