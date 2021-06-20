package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeBool interface {
	Execute(df *dataframe.DataFrame) bool
}

type PipelineBool struct {
	Pipes Pipeline
	Tail  PipeBool
}

func (line PipelineBool) Execute(frame *dataframe.DataFrame) bool {
	frame = line.Pipes.Execute(frame)
	return line.Tail.Execute(frame)
}
