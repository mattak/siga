package cmd

import (
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	FilterNotOrCmd = &cobra.Command{
		Use:     "nor [NUMBER] [COLUMN_NAME]+",
		Aliases: []string{"no"},
		Short:   "Filter sequence by NOT operation",
		Long:    "Filter sequence by NOT operation",
		Example: `
  siga filter nor 1 column1
  siga filter nor NaN column1 column2
`,
		Run: runCommandFilterNotOr,
	}
)

func init() {
}

func runCommandFilterNotOr(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("[NUMBER] [COLUMN_NAME]+ should be declared")
	}
	targetValue := toolkit.ParseFloat64(args[0])
	columnNames := args[1:]

	df := toolkit.ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(columnNames)
	indexes := matrix.FilterIndexByNotOr(targetValue)
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
