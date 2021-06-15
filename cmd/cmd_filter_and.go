package cmd

import (
	"github.com/mattak/siga/pkg"
	"github.com/spf13/cobra"
	"log"
)

var (
	FilterAndCmd = &cobra.Command{
		Use:     "and [NUMBER] [COLUMN_NAME]+",
		Aliases: []string{"a"},
		Short:   "Filter sequence by AND operation",
		Long:    "Filter sequence by AND operation",
		Example: `
  siga filter and 1 column1
  siga filter and 1 column1 column2
`,
		Run: runCommandFilterAnd,
	}
)

func init() {
}

func runCommandFilterAnd(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("[NUMBER] [COLUMN_NAME]+ should be declared")
	}
	targetValue := pkg.ParseFloat64(args[0])
	columnNames := args[1:]

	df := pkg.ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(columnNames)
	indexes := matrix.FilterIndexByAnd(targetValue)
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
