package pipeline

import "github.com/mattak/siga/pkg/dataframe"

type PipeFloat64 interface {
	Execute(df *dataframe.DataFrame) float64
}

type PipelineFloat64 struct {
	Pipe Pipe
	Tail PipeFloat64
}

func (line PipelineFloat64) Execute(frame *dataframe.DataFrame) float64 {
	frame = line.Pipe.Execute(frame)
	return line.Tail.Execute(frame)
}
