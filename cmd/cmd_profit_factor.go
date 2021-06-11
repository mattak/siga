package cmd

import (
	"github.com/mattak/siga/toolkit"
	"github.com/spf13/cobra"
	"log"
)

var (
	ProfitFactorCmd = &cobra.Command{
		Use:     "profit_factor [COLUMN_NAME]+",
		Aliases: []string{"pf"},

		Short: "Calculate profit factor",
		Long:  "Calculate profit factor",
		Example: `
  siga pf column1 < sample.tsv
  siga pf column1 column2 < sample.tsv
`,
		Run: runCommandProfitFactor,
	}
)

func init() {
}

func runCommandProfitFactor(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("More than one COLUMN_NAME should be declared")
	}

	df := toolkit.ReadDataFrameByStdinTsv()
	result_vector := make(toolkit.Vector, len(args))

	for i := 0; i < len(args); i++ {
		vector, err := df.ExtractColumn(args[i])
		if err != nil {
			log.Fatalf("COLUMN_NAME or NUMBER required: %s", args[i])
		}

		result_vector[i] = vector.ProfitFactor()
	}

	result_vector.PrintTsv(IsPreciseOutput)
}
