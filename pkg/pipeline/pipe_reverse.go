package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type ReversePipe struct {
}

func (pipe ReversePipe) Execute(df *dataframe.DataFrame) *dataframe.DataFrame {
	df.Reverse()
	return df
}
