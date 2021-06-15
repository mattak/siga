package cmd

import (
	"github.com/mattak/siga/pkg"
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
	df := pkg.ReadDataFrameByStdinTsv()
	df.Reverse()
	df.PrintTsv(IsPreciseOutput)
}
