package pipeline

import (
	"github.com/mattak/siga/pkg/dataframe"
)

type CountCommandOption struct {
}

type CountCommandPipeOutput struct {
	DataFrame *dataframe.DataFrame
	Option    CountCommandOption
}

func (c CobraCommandInput) CreateCountCommandOption() CountCommandOption {
	return CountCommandOption{ }
}

func (option CountCommandOption) CreatePipeInt(df *dataframe.DataFrame) PipeInt {
	return CountCommandPipeOutput{
		DataFrame: df,
		Option:    option,
	}
}

func (pipe CountCommandPipeOutput) Execute() int {
	return len(pipe.DataFrame.Labels)
}
