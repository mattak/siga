package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
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
	df := dataframe.ReadDataFrameByStdinTsv()
	if len(args) < 1 {
		log.Fatal("SIZE should be declared")
	}
	size := util.ParseInt(args[0])
	df.Take(size)
	df.PrintTsv(IsPreciseOutput)
}
