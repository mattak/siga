package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	SelectCmd = &cobra.Command{
		Use:     "select [COLUMN_NAME]+",
		Aliases: []string{"sl"},

		Short: "select column",
		Long:  "Select column",
		Example: `
Select column of open, close, volume:
  siga sl open close volume < sample.tsv
`,
		Run: runCommandSelect,
	}
)

func init() {
}

func runCommandSelect(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("COLUMN_NAME should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	err := df.SelectColumn(args...)
	if err != nil {
		log.Fatal(err)
	}
	df.PrintTsv(IsPreciseOutput)
}
