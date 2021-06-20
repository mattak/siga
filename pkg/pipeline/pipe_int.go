package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeInt interface {
	Execute(df *dataframe.DataFrame) int
}

type PipelineInt struct {
	Pipe Pipe
	Tail PipeInt
}

func (line PipelineInt) Execute(frame *dataframe.DataFrame) int {
	frame = line.Pipe.Execute(frame)
	return line.Tail.Execute(frame)
}
