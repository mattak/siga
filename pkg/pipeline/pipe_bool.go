package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeBool interface {
	Execute() bool
}

type PipeBoolCreator interface {
	CreatePipeBool(df *dataframe.DataFrame) PipeBool
}

type PipelineBool struct {
	Pipes   Pipeline
	Creator PipeBoolCreator
}

func (line PipelineBool) Execute(frame *dataframe.DataFrame) bool {
	frame = line.Pipes.ExecutePipes(frame)
	return line.Creator.CreatePipeBool(frame).Execute()
}
