package cmd

import (
	"fmt"
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	DollarCostAverageCmd = &cobra.Command{
		Use:     "dollar_cost_average [COLUMN_NAME]+",
		Aliases: []string{"dollar_cost", "dca", "dc"},

		Short: "DollarCost",
		Long:  "Calculate DollarCost average prices",
		Example: `
  siga dc close < sample.tsv
`,
		Run: runCommandDollarCost,
	}
)

func init() {
	DollarCostAverageCmd.Flags().StringVarP(&label, "label", "l", "", "overwrite label name")
}

func runCommandDollarCost(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than one COLUMN_NAME or NUMBER should be declared")
	}

	df := toolkit.ReadDataFrameByStdinTsv()

	results := make([]toolkit.Vector, len(args))
	for i := 0; i < len(args); i++ {
		vector, err := df.ExtractColumn(args[i])
		if err == nil {
			results[i] = vector
		} else {
			log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
		}

		results[i] = vector.DollarCostAverage()
	}

	for i := 0; i < len(args); i++ {
		columnName := label
		if label == "" {
			columnName = fmt.Sprintf("dca_%s", args[i])
		}

		err := df.AddColumn(columnName, results[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	df.PrintTsv(IsPreciseOutput)
}
