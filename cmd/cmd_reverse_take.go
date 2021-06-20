package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	ReverseTakeCmd = &cobra.Command{
		Use:     "reverse_take [SIZE]",
		Aliases: []string{"retake"},

		Short: "ReverseTake data",
		Long:  "ReverseTake data",
		Example: `
  siga retake 10 < sample.tsv
`,
		Run: runCommandReverseTake,
	}
)

func init() {
}

func runCommandReverseTake(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	creator := pipeline.CobraCommandInput{cmd, args}.CreateReverseTakeCommandOption(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.CreatePipe(df).Execute()

	df.PrintTsv(IsPreciseOutput)
}
