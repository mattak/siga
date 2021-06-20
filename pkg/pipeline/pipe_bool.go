package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeBool interface {
	Execute(df *dataframe.DataFrame) bool
}

type PipelineBool struct {
	Pipe Pipe
	Tail PipeBool
}

func (line PipelineBool) Execute(frame *dataframe.DataFrame) bool {
	frame = line.Pipe.Execute(frame)
	return line.Tail.Execute(frame)
}
