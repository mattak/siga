package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type ReverseCommandOption struct {
}

type ReverseCommandPipe struct {
	DataFrame *dataframe.DataFrame
}

func (option ReverseCommandOption) CreatePipe(df *dataframe.DataFrame) Pipe {
	return ReverseCommandPipe{
		DataFrame: df,
	}
}

func (pipe ReverseCommandPipe) Execute() *dataframe.DataFrame {
	pipe.DataFrame.Reverse()
	return pipe.DataFrame
}