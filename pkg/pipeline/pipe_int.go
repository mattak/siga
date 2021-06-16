package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeInt interface {
	Execute() int
}

type PipeIntCreator interface {
	CreatePipeInt(df *dataframe.DataFrame) PipeInt
}

type PipelineInt struct {
	Pipes   Pipeline
	Creator PipeIntCreator
}

func (line PipelineInt) Execute(frame *dataframe.DataFrame) int {
	frame = line.Pipes.ExecutePipes(frame)
	return line.Creator.CreatePipeInt(frame).Execute()
}
