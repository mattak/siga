package cmd

import (
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	FilterNotAndCmd = &cobra.Command{
		Use:     "nand [COLUMN_NAME|NUMBER] [COLUMN_NAME|NUMBER]+",
		Aliases: []string{"na"},
		Short:   "Filter sequence by NOT operation",
		Long:    "Filter sequence by NOT operation",
		Example: `
  siga filter nand 1 column1
  siga filter nand NaN column1 column2
`,
		Run: runCommandFilterNotAnd,
	}
)

func init() {
}

func runCommandFilterNotAnd(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("[NUMBER] [COLUMN_NAME]+ should be declared")
	}
	targetValue := toolkit.ParseFloat64(args[0])
	columnNames := args[1:]

	df := toolkit.ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(columnNames)
	indexes := matrix.FilterIndexByNotAnd(targetValue)
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
