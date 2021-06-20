package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeInt interface {
	Execute(df *dataframe.DataFrame) int
}

type PipelineInt struct {
	Pipes Pipeline
	Tail  PipeInt
}

func (line PipelineInt) Execute(frame *dataframe.DataFrame) int {
	frame = line.Pipes.Execute(frame)
	return line.Tail.Execute(frame)
}
