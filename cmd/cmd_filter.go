package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	FilterCmd = &cobra.Command{
		Use:     "filter [COLUMN_NAME] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"fl"},
		Short:   "Filter sequence by value",
		Long:    "Filter sequence by value",
		Example: `
  siga filter column1 1
  siga filter column1 column2
`,
		Run: runCommandFilter,
	}
)

func init() {
}

func runCommandFilter(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(args)
	indexes := matrix.FilterIndex()
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
