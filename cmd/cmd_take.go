package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	TakeCmd = &cobra.Command{
		Use:     "take [SIZE]",
		Aliases: []string{"tk"},

		Short: "Take head data",
		Long:  "Take head data",
		Example: `
  siga take 10 < sample.tsv
`,
		Run: runCommandTake,
	}
)

func init() {
}

func runCommandTake(cmd *cobra.Command, args []string) {
	outputOption := pipeline.OutputOption{ColumnName: label}
	creator := pipeline.CobraCommandInput{cmd, args}.CreateTakePipe(outputOption)
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.Execute(df)

	df.PrintTsv(IsPreciseOutput)
}
