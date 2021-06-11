package cmd

import (
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	FilterNotCmd = &cobra.Command{
		Use:     "filter_not [COLUMN_NAME] [COLUMN_NAME|NUMBER]",
		Aliases: []string{"fn"},
		Short:   "FilterNot sequence by value",
		Long:    "FilterNot sequence by value",
		Example: `
  siga filter_not column1 1
  siga filter_not column1 column2
`,
		Run: runCommandFilterNot,
	}
)

func init() {
}

func runCommandFilterNot(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("Two COLUMN_NAME or NUMBER should be declared")
	}

	df := toolkit.ReadDataFrameByStdinTsv()

	matrix := df.ExtractMatrixByColumnNameOrValue(args)
	indexes := matrix.FilterNotIndex()
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
