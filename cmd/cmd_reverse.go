package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/pipeline"
	"github.com/spf13/cobra"
)

var (
	ReverseCmd = &cobra.Command{
		Use:     "reverse",
		Aliases: []string{"re", "rev"},

		Short: "Identity mapping",
		Long:  "Identity mapping",
		Example: `
  siga re < sample.tsv
`,
		Run: runCommandReverse,
	}
)

func init() {
}

func runCommandReverse(cmd *cobra.Command, args []string) {
	creator := pipeline.ReverseCommandOption{}
	df := dataframe.ReadDataFrameByStdinTsv()
	df = creator.CreatePipe(df).Execute()
	df.PrintTsv(IsPreciseOutput)
}
