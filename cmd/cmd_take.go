package cmd

import (
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
	df := ReadDataFrameByStdinTsv()
	if len(args) < 1 {
		log.Fatal("SIZE should be declared")
	}
	size := ParseInt(args[0])
	df.Take(size)
	df.PrintTsv(IsPreciseOutput)
}
