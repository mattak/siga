package cmd

import (
	"github.com/mattak/siga/pkg/dataframe"
	"github.com/mattak/siga/pkg/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	FilterOrCmd = &cobra.Command{
		Use:     "or [NUMBER] [COLUMN_NAME]+",
		Aliases: []string{"o"},
		Short:   "Filter or sequence by OR operation",
		Long:    "Filter or sequence by OR operation",
		Example: `
  siga filter or 1 column1
  siga filter or NaN column1 column2
`,
		Run: runCommandFilterOr,
	}
)

func init() {
}

func runCommandFilterOr(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatal("[NUMBER] [COLUMN_NAME]+ should be declared")
	}
	targetValue := util.ParseFloat64(args[0])
	columnNames := args[1:]

	df := dataframe.ReadDataFrameByStdinTsv()
	matrix := df.ExtractMatrixByColumnNameOrValue(columnNames)
	indexes := matrix.FilterIndexByOr(targetValue)
	df.SelectRecords(indexes...)
	df.PrintTsv(IsPreciseOutput)
}
