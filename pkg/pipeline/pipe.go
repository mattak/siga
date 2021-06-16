package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type Pipe interface {
	Execute() *dataframe.DataFrame
}

type PipeCreator interface {
	CreatePipe(df *dataframe.DataFrame) Pipe
}

type Pipeline []PipeCreator

func (creators Pipeline) ExecutePipes(df *dataframe.DataFrame) *dataframe.DataFrame {
	for i := 0; i < len(creators); i++ {
		df = creators[i].CreatePipe(df).Execute()
	}
	return df
}
